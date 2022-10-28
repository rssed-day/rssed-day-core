package plugins

import (
	"github.com/rssed-day/rssed-day-core/models"
)

// Action -
type Action interface {
	AddObject(obj models.Object)
	AddError(err error)
}

// action -
type action struct {
	runner  Runner
	objects chan<- models.Object
	err     error
}

// NewAction -
func NewAction(runner Runner, objects chan<- models.Object) Action {
	return &action{
		runner:  runner,
		objects: objects,
	}
}

// AddObject -
func (c *action) AddObject(obj models.Object) {
	if obj == nil {
		return
	}
	c.objects <- obj
}

// AddError -
func (c *action) AddError(err error) {
	if err == nil {
		return
	}
	c.err = err
	c.runner.Logger().Error(err)
}
