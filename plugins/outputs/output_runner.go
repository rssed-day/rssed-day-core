package outputs

import (
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// OutputConfig -
type OutputConfig struct {
	Uuid string `yaml:"uuid"`
	Name string `yaml:"name"`

	Interval    time.Duration `yaml:"interval"`
	BufferLimit int           `yaml:"buffer_limit"`
	BufferBatch int           `yaml:"buffer_batch"`

	Args map[string]interface{} `yaml:"args"`
}

// OutputRunner -
type OutputRunner struct {
	Output OutputPlugin
	Config *OutputConfig

	OutputCount    int64
	OutputDuration time.Duration

	buffer *Buffer

	logger logrus.FieldLogger
}

// NewOutputRunner -
func NewOutputRunner(output OutputPlugin, config *OutputConfig) *OutputRunner {
	// TODO: set output logger
	return &OutputRunner{
		Output: output,
		Config: config,
		logger: logrus.StandardLogger(),
	}
}

// Init -
func (r *OutputRunner) Init() error {
	if i, ok := r.Output.(plugins.Initializer); ok {
		if err := i.Init(); err != nil {
			return err
		}
	}
	return nil
}

// FanOut -
func (r *OutputRunner) FanOut(in ...models.Object) error {
	// TODO: duration and count
	return r.Output.FanOut(in...)
}

// FanOutAll -
func (r *OutputRunner) FanOutAll() error {
	// TODO
	return nil
}

// FanOutBatch -
func (r *OutputRunner) FanOutBatch() error {
	// TODO
	return nil
}

// Connect -
func (r *OutputRunner) Connect() error {
	return r.Output.Connect()
}

// Close -
func (r *OutputRunner) Close() error {
	return r.Output.Close()
}

// Logger -
func (r *OutputRunner) Logger() logrus.FieldLogger {
	return r.logger
}

// OutputRunners -
type OutputRunners []*OutputRunner

// Buffer -
type Buffer struct {
	mutex   sync.Mutex
	objects []models.Object

	BufferLimit int
	BufferBatch int
}
