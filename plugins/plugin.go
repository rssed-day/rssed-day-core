package plugins

// Describer -
type Describer interface {
	Description() string
}

// Plugin -
type Plugin interface {
	Describer
}
