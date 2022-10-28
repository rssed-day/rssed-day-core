package inputs

import (
	"github.com/rssed-day/rssed-day-core/plugins"
	"github.com/sirupsen/logrus"
	"time"
)

// InputConfig -
type InputConfig struct {
	Uuid string `yaml:"uuid"`
	Name string `yaml:"name"`

	Cron string `yaml:"cron"`

	Args map[string]interface{} `yaml:"args"`
}

// InputRunner -
type InputRunner struct {
	Input  InputPlugin
	Config *InputConfig

	InputCount    int64
	InputDuration time.Duration

	logger logrus.FieldLogger
}

// NewInputRunner -
func NewInputRunner(input InputPlugin, config *InputConfig) *InputRunner {
	// TODO: set input logger
	return &InputRunner{
		Input:  input,
		Config: config,
		logger: logrus.StandardLogger(),
	}
}

// Init -
func (r *InputRunner) Init() error {
	if i, ok := r.Input.(plugins.Initializer); ok {
		if err := i.Init(); err != nil {
			return err
		}
	}
	return nil
}

// FanIn -
func (r *InputRunner) FanIn(out plugins.Action) error {
	// TODO: duration and count
	return r.Input.FanIn(out)
}

// Logger -
func (r *InputRunner) Logger() logrus.FieldLogger {
	return r.logger
}

// InputRunners -
type InputRunners []*InputRunner
