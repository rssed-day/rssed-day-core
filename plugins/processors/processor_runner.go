package processors

import (
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
	"github.com/sirupsen/logrus"
	"time"
)

// ProcessorConfig -
type ProcessorConfig struct {
	Uuid string `yaml:"uuid"`
	Name string `yaml:"name"`

	Order int `yaml:"order"`

	Args map[string]interface{} `yaml:"args"`
}

// ProcessorRunner -
type ProcessorRunner struct {
	Processor ProcessorPlugin
	Config    *ProcessorConfig

	ProcessorCount    int64
	ProcessorDuration time.Duration

	logger logrus.FieldLogger
}

// NewProcessorRunner -
func NewProcessorRunner(processor ProcessorPlugin, config *ProcessorConfig) *ProcessorRunner {
	// TODO: set ProcessorPlugin logger
	return &ProcessorRunner{
		Processor: processor,
		Config:    config,
		logger:    logrus.StandardLogger(),
	}
}

// Init -
func (r *ProcessorRunner) Init() error {
	if i, ok := r.Processor.(plugins.Initializer); ok {
		if err := i.Init(); err != nil {
			return err
		}
	}
	return nil
}

// Process -
func (r *ProcessorRunner) Process(in models.Object, out plugins.Action) error {
	// TODO: duration and count
	return r.Processor.Process(in, out)
}

// Logger -
func (r *ProcessorRunner) Logger() logrus.FieldLogger {
	return r.logger
}

// ProcessorRunners -
type ProcessorRunners []*ProcessorRunner

// Len -
func (r ProcessorRunners) Len() int { return len(r) }

// Swap -
func (r ProcessorRunners) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

// Less -
func (r ProcessorRunners) Less(i, j int) bool { return r[i].Config.Order < r[j].Config.Order }
