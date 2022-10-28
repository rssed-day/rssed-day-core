package parsers

import (
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
)

// Parser -
type Parser interface {
	plugins.Describer

	Parse([]byte) (models.Object, error)
	ParseBatch([]byte) ([]models.Object, error)
}

// ParserPlugin -
type ParserPlugin interface {
	SetParser(parser Parser)
}

// ParserFactory -
type ParserFactory func() Parser

// ParserFactories -
var ParserFactories = map[string]ParserFactory{}

// Register -
func Register(name string, factory ParserFactory) {
	ParserFactories[name] = factory
}
