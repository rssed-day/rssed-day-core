package serializers

import (
	"github.com/rssed-day/rssed-day-core/models"
	"github.com/rssed-day/rssed-day-core/plugins"
)

// Serializer -
type Serializer interface {
	plugins.Describer

	Serialize(models.Object) []byte
	SerializeBatch([]models.Object) []byte
}

// SerializerPlugin -
type SerializerPlugin interface {
	SetSerializer(serializer Serializer)
}

// SerializerFactory -
type SerializerFactory func() Serializer

// SerializerFactories -
var SerializerFactories = map[string]SerializerFactory{}

// Register -
func Register(name string, factory SerializerFactory) {
	SerializerFactories[name] = factory
}
