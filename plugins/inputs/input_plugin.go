package inputs

import (
	"github.com/rssed-day/rssed-day-core/plugins"
)

// InputPlugin -
type InputPlugin interface {
	plugins.Describer

	FanIn(out plugins.Action) error
}

// InputFactory -
type InputFactory func() InputPlugin

// InputFactories -
var InputFactories = map[string]InputFactory{}

// Register -
func Register(name string, factory InputFactory) {
	InputFactories[name] = factory
}
