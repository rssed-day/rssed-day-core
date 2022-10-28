package models

// Metadata -
type Metadata struct {
	Uuid     string
	Name     string
	Protocol string
}

// Object -
type Object interface {
	Meta() Metadata
	Raw() []byte
}

// object -
type object struct {
	meta Metadata
	raw  []byte
}

// NewObject -
func NewObject(metadata Metadata, raw []byte) Object {
	return &object{
		meta: metadata,
		raw:  raw,
	}
}

// Meta -
func (o *object) Meta() Metadata {
	return o.meta
}

// Raw -
func (o *object) Raw() []byte {
	return o.raw
}
