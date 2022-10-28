package inputs

import "github.com/rssed-day/rssed-day-core/models"

// InputActions -
//
// ┌───────┐
// │ InputPlugin │───┐
// └───────┘   │
// ┌───────┐   │     ______
// │ InputPlugin │───┼──▶ ()_____)
// └───────┘   │
// ┌───────┐   │
// │ InputPlugin │───┘
// └───────┘
type InputActions struct {
	Dst    chan<- models.Object
	Inputs []*InputRunner
}
