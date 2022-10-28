package processors

import "github.com/rssed-day/rssed-day-core/models"

// ProcessorActions -
//  ______     ┌───────────┐     ______
// ()_____)──▶ │ ProcessorPlugin │──▶ ()_____)
//             └───────────┘
type ProcessorActions struct {
	Src       <-chan models.Object
	Dst       chan<- models.Object
	Processor *ProcessorRunner
}
