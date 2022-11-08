package processors

import "github.com/rssed-day/rssed-day-core/models"

// PreProcessorActions -
//  ______     ┌───────────┐     ______
// ()_____)──▶ │ Processor │──▶ ()_____)
//             └───────────┘
type PreProcessorActions struct {
	Src       <-chan models.Object
	Dst       chan<- models.Object
	Processor *ProcessorRunner
}

// PostProcessorActions -
//  ______     ┌───────────┐     ______
// ()_____)──▶ │ Processor │──▶ ()_____)
//             └───────────┘
type PostProcessorActions struct {
	Src       <-chan models.Object
	Dst       chan<- models.Object
	Processor *ProcessorRunner
}
