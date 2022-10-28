package aggregators

import (
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// AggregatorConfig -
type AggregatorConfig struct {
	Uuid string `yaml:"uuid"`
	Name string `yaml:"name"`

	Args map[string]interface{} `yaml:"args"`
}

// AggregatorRunner -
type AggregatorRunner struct {
	sync.Mutex

	Aggregator AggregatorPlugin
	Config     *AggregatorConfig

	AggregatorCount    int64
	AggregatorDuration time.Duration

	logger logrus.FieldLogger
}

// NewAggregatorRunner -
func NewAggregatorRunner(aggregator AggregatorPlugin, config *AggregatorConfig) *AggregatorRunner {
	// TODO: set aggregator logger
	return &AggregatorRunner{
		Aggregator: aggregator,
		Config:     config,
		logger:     logrus.StandardLogger(),
	}
}

// Init -
func (r *AggregatorRunner) Init() error {
	if i, ok := r.Aggregator.(plugins.Initializer); ok {
		if err := i.Init(); err != nil {
			return err
		}
	}
	return nil
}

// Gather -
func (r *AggregatorRunner) Gather(in models.Object) {
	// TODO: duration and count
	r.Lock()
	defer r.Unlock()
	r.Aggregator.Gather(in)
	return
}

// Flush -
func (r *AggregatorRunner) Flush(out plugins.Action) {
	// TODO: duration and count
	r.Lock()
	defer r.Unlock()
	r.Aggregator.Flush(out)
	r.Aggregator.Reset()
	return
}

// Logger -
func (r *AggregatorRunner) Logger() logrus.FieldLogger {
	return r.logger
}

// AggregatorRunners -
type AggregatorRunners []*AggregatorRunner
