package context

import (
	"fmt"
	"github.com/rssed-day/rssed-day-core/plugins/aggregators"
	"github.com/rssed-day/rssed-day-core/plugins/inputs"
	"github.com/rssed-day/rssed-day-core/plugins/outputs"
	"github.com/rssed-day/rssed-day-core/plugins/parsers"
	"github.com/rssed-day/rssed-day-core/plugins/processors"
	"github.com/rssed-day/rssed-day-core/plugins/serializers"
	"gopkg.in/yaml.v3"
)

// Context -
type Context struct {
	Config            *Config
	InputRunners      inputs.InputRunners
	ProcessorRunners  processors.ProcessorRunners
	AggregatorRunners aggregators.AggregatorRunners
	OutputRunners     outputs.OutputRunners
}

// NewContext -
func NewContext(config *Config) (*Context, error) {
	context := &Context{
		Config: config,
	}
	if err := context.setInputRunners(config.InputConfigs); err != nil {
		return nil, err
	}
	if err := context.setProcessorRunners(config.ProcessorConfigs); err != nil {
		return nil, err
	}
	if err := context.setAggregatorRunners(config.AggregatorConfigs); err != nil {
		return nil, err
	}
	if err := context.setOutputRunners(config.OutputConfigs); err != nil {
		return nil, err
	}
	return context, nil
}

// setInputRunners -
func (c *Context) setInputRunners(configs []inputs.InputConfig) error {
	for _, config := range configs {
		factory, ok := inputs.InputFactories[config.Name]
		if !ok {
			return fmt.Errorf("no such %s input plugin", config.Name)
		}
		plugin := factory()

		if _, ok := plugin.(parsers.ParserPlugin); ok {
			// TODO: set parser
		}

		runner := inputs.NewInputRunner(plugin, &config)
		args, err := yaml.Marshal(config.Args)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(args, runner.Input)
		if err != nil {
			return err
		}
		c.InputRunners = append(c.InputRunners, runner)
	}
	return nil
}

// setProcessorRunners -
func (c *Context) setProcessorRunners(configs []processors.ProcessorConfig) error {
	for _, config := range configs {
		factory, ok := processors.ProcessorFactories[config.Name]
		if !ok {
			return fmt.Errorf("no such %s processor plugin", config.Name)
		}
		plugin := factory()

		if _, ok := plugin.(parsers.ParserPlugin); ok {
			// TODO: set parser
		}

		if _, ok := plugin.(serializers.SerializerPlugin); ok {
			// TODO: set serializer
		}

		runner := processors.NewProcessorRunner(plugin, &config)
		args, err := yaml.Marshal(config.Args)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(args, runner.Processor)
		if err != nil {
			return err
		}
		c.ProcessorRunners = append(c.ProcessorRunners, runner)
	}
	return nil
}

// setAggregatorRunners -
func (c *Context) setAggregatorRunners(configs []aggregators.AggregatorConfig) error {
	for _, config := range configs {
		factory, ok := aggregators.AggregatorFactories[config.Name]
		if !ok {
			return fmt.Errorf("no such %s aggregator plugin", config.Name)
		}
		plugin := factory()

		if _, ok := plugin.(parsers.ParserPlugin); ok {
			// TODO: set parser
		}

		if _, ok := plugin.(serializers.SerializerPlugin); ok {
			// TODO: set serializer
		}

		runner := aggregators.NewAggregatorRunner(plugin, &config)
		args, err := yaml.Marshal(config.Args)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(args, runner.Aggregator)
		if err != nil {
			return err
		}
		c.AggregatorRunners = append(c.AggregatorRunners, runner)
	}
	return nil
}

// setOutputRunners -
func (c *Context) setOutputRunners(configs []outputs.OutputConfig) error {
	for _, config := range configs {
		factory, ok := outputs.OutputFactories[config.Name]
		if !ok {
			return fmt.Errorf("no such %s output plugin", config.Name)
		}
		plugin := factory()

		if _, ok := plugin.(serializers.SerializerPlugin); ok {
			// TODO: set serializer
		}

		runner := outputs.NewOutputRunner(plugin, &config)
		args, err := yaml.Marshal(config.Args)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(args, runner.Output)
		if err != nil {
			return err
		}
		c.OutputRunners = append(c.OutputRunners, runner)
	}
	return nil
}
