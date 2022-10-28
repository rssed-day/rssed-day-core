package processors

import (
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
)

// ProcessorPlugin -
type ProcessorPlugin interface {
	plugins.Describer

	Process(in models.Object, out plugins.Action) error
	process(...models.Object) []models.Object
}

// ProcessorFactory -
type ProcessorFactory func() ProcessorPlugin

// ProcessorFactories -
var ProcessorFactories = map[string]ProcessorFactory{}

// Register -
func Register(name string, factory ProcessorFactory) {
	ProcessorFactories[name] = factory
}
