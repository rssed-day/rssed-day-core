package aggregators

import (
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
)

// AggregatorPlugin -
type AggregatorPlugin interface {
	plugins.Describer

	Gather(in models.Object)
	Flush(out plugins.Action)
	Reset()
}

// AggregatorFactory -
type AggregatorFactory func() AggregatorPlugin

// AggregatorFactories -
var AggregatorFactories = map[string]AggregatorFactory{}

// Register -
func Register(name string, factory AggregatorFactory) {
	AggregatorFactories[name] = factory
}
