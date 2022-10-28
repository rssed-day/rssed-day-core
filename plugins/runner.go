package plugins

import "github.com/sirupsen/logrus"

// Initializer -
type Initializer interface {
	Init() error
}

// Runner -
type Runner interface {
	Initializer
	Logger() logrus.FieldLogger
}
