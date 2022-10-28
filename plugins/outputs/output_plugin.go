package outputs

import (
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
)

// OutputPlugin -
type OutputPlugin interface {
	plugins.Describer

	FanOut(in ...models.Object) error
	Connect() error
	Close() error
}

// OutputFactory -
type OutputFactory func() OutputPlugin

// OutputFactories -
var OutputFactories = map[string]OutputFactory{}

// Register -
func Register(name string, factory OutputFactory) {
	OutputFactories[name] = factory
}
