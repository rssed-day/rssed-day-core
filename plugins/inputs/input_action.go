package inputs

import "github.com/rssed-day/rssed-day-core/models"

// InputActions -
//
// ┌───────┐
// │ Input │───┐
// └───────┘   │
// ┌───────┐   │     ______
// │ Input │───┼──▶ ()_____)
// └───────┘   │
// ┌───────┐   │
// │ Input │───┘
// └───────┘
type InputActions struct {
	Dst    chan<- models.Object
	Inputs []*InputRunner
}
