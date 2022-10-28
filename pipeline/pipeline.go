package pipeline

import (
	"context"
	"github.com/rssed-day/rssed-day-core/constants"
	context2 "github.com/rssed-day/rssed-day-core/context"
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
	"github.com/rssed-day/rssed-day-core/plugins/aggregators"
	"github.com/rssed-day/rssed-day-core/plugins/inputs"
	"github.com/rssed-day/rssed-day-core/plugins/outputs"
	"github.com/rssed-day/rssed-day-core/plugins/processors"
	"sort"
	"sync"
)

// Pipeline -
type Pipeline struct {
	Context *context2.Context
}

// NewPipeline -
func NewPipeline(context *context2.Context) *Pipeline {
	p := &Pipeline{
		Context: context,
	}
	return p
}

// Pipe -
func (p *Pipeline) Pipe(ctx context.Context) error {
	var err error

	if err := p.constructPlugins(); err != nil {
		return err
	}

	next, outputUnit, err := p.assembleOutputs(p.Context.OutputRunners)
	if err != nil {
		return err
	}

	var aggregatorUnit *aggregators.AggregatorActions
	if len(p.Context.AggregatorRunners) != 0 {
		next, aggregatorUnit, err = p.assembleAggregators(next, p.Context.AggregatorRunners)
		if err != nil {
			return err
		}
	}

	var processorUnits []*processors.ProcessorActions
	if len(p.Context.ProcessorRunners) != 0 {
		next, processorUnits, err = p.assembleProcessors(next, p.Context.ProcessorRunners)
		if err != nil {
			return err
		}
	}

	inputUnit, err := p.assembleInputs(next, p.Context.InputRunners)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		p.pipeOutputs(ctx, outputUnit)
	}()

	if aggregatorUnit != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.pipeAggregators(ctx, aggregatorUnit)
		}()
	}

	if processorUnits != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.pipeProcessors(ctx, processorUnits)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		p.pipeInputs(ctx, inputUnit)
	}()

	wg.Wait()

	if err := p.destructPlugins(); err != nil {
		return err
	}

	return nil
}

// constructPlugins -
func (p *Pipeline) constructPlugins() error {
	for _, plugin := range p.Context.InputRunners {
		if err := plugin.Init(); err != nil {
			return err
		}
	}
	for _, plugin := range p.Context.ProcessorRunners {
		if err := plugin.Init(); err != nil {
			return err
		}
	}
	for _, plugin := range p.Context.AggregatorRunners {
		if err := plugin.Init(); err != nil {
			return err
		}
	}
	for _, plugin := range p.Context.OutputRunners {
		if err := plugin.Init(); err != nil {
			return err
		}
	}
	return nil
}

// destructPlugins -
func (p *Pipeline) destructPlugins() error {
	return nil
}

// assembleInputs -
func (p *Pipeline) assembleInputs(dst chan<- models.Object, inputRunners []*inputs.InputRunner) (
	*inputs.InputActions, error) {
	unit := &inputs.InputActions{
		Dst:    dst,
		Inputs: inputRunners,
	}
	return unit, nil
}

// pipeInputs -
func (p *Pipeline) pipeInputs(ctx context.Context, unit *inputs.InputActions) {
	var wg sync.WaitGroup

	for _, input := range unit.Inputs {
		act := plugins.NewAction(input, unit.Dst)

		// TODO: input with cron

		wg.Add(1)
		go func(runner *inputs.InputRunner) {
			defer wg.Done()
			if err := runner.FanIn(act); err != nil {
				act.AddError(err)
			}
		}(input)
	}

	wg.Wait()

	close(unit.Dst)
}

// assembleProcessors -
func (p *Pipeline) assembleProcessors(dst chan<- models.Object, processorRunners processors.ProcessorRunners) (
	chan<- models.Object, []*processors.ProcessorActions, error) {
	var units []*processors.ProcessorActions

	sort.SliceStable(processorRunners, func(i, j int) bool {
		return processorRunners[i].Config.Order > processorRunners[j].Config.Order
	})

	var src chan models.Object
	for _, processor := range processorRunners {
		src = make(chan models.Object, constants.ObjectChannelSize)

		unit := &processors.ProcessorActions{
			Src:       src,
			Dst:       dst,
			Processor: processor,
		}
		units = append(units, unit)

		dst = src
	}

	return src, units, nil
}

// pipeProcessors -
func (p *Pipeline) pipeProcessors(ctx context.Context, units []*processors.ProcessorActions) {
	var wg sync.WaitGroup
	for _, unit := range units {
		wg.Add(1)
		go func(unit *processors.ProcessorActions) {
			defer wg.Done()

			act := plugins.NewAction(unit.Processor, unit.Dst)
			for obj := range unit.Src {
				if err := unit.Processor.Process(obj, act); err != nil {
					act.AddError(err)
				}
			}
			close(unit.Dst)
		}(unit)
	}
	wg.Wait()
}

// assembleAggregators -
func (p *Pipeline) assembleAggregators(dst chan<- models.Object, aggregatorRunners []*aggregators.AggregatorRunner) (
	chan<- models.Object, *aggregators.AggregatorActions, error) {
	src := make(chan models.Object, constants.ObjectChannelSize)
	unit := &aggregators.AggregatorActions{
		Src:         src,
		Dst:         dst,
		Aggregators: aggregatorRunners,
	}
	return src, unit, nil
}

// pipeAggregators -
func (p *Pipeline) pipeAggregators(ctx context.Context, unit *aggregators.AggregatorActions) {
	var wg sync.WaitGroup

	// gather iteratively by objects from src
	wg.Add(1)
	go func() {
		defer wg.Done()

		for obj := range unit.Src {
			for _, aggregator := range unit.Aggregators {
				aggregator.Gather(obj)
			}
		}
	}()

	// flush concurrently by aggregators to dst
	for _, aggregator := range unit.Aggregators {
		wg.Add(1)
		go func(aggregator *aggregators.AggregatorRunner) {
			defer wg.Done()

			act := plugins.NewAction(aggregator, unit.Dst)
			aggregator.Flush(act)
		}(aggregator)
	}

	wg.Wait()

	close(unit.Dst)
}

// assembleOutputs -
func (p *Pipeline) assembleOutputs(outputRunners []*outputs.OutputRunner) (
	chan<- models.Object, *outputs.OutputActions, error) {
	src := make(chan models.Object, constants.ObjectChannelSize)

	for _, output := range outputRunners {
		if err := output.Connect(); err != nil {
			return src, nil, err
		}
	}

	unit := &outputs.OutputActions{
		Src:     src,
		Outputs: outputRunners,
	}
	return src, unit, nil
}

// pipeOutputs -
func (p *Pipeline) pipeOutputs(ctx context.Context, unit *outputs.OutputActions) {
	var wg sync.WaitGroup

	// TODO: output with buffer

	for _, output := range unit.Outputs {
		wg.Add(1)
		go func(output *outputs.OutputRunner) {
			defer wg.Done()
			for obj := range unit.Src {
				if err := output.FanOut(obj); err != nil {
					// TODO: log
				}
			}
		}(output)
	}

	wg.Wait()

	for _, output := range unit.Outputs {
		if err := output.Close(); err != nil {
			// TODO: log
		}
	}
}
