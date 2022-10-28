package aggregators

import "github.com/rssed-day/rssed-day-core/models"

// AggregatorActions -
//                 ┌────────────┐
//            ┌──▶ │ AggregatorPlugin │───┐
//            │    └────────────┘   │
//  ______    │    ┌────────────┐   │     ______
// ()_____)───┼──▶ │ AggregatorPlugin │───┼──▶ ()_____)
//            │    └────────────┘   │
//            │    ┌────────────┐   │
//            └──▶ │ AggregatorPlugin │───┘
//                 └────────────┘
type AggregatorActions struct {
	Src         <-chan models.Object
	Dst         chan<- models.Object
	Aggregators []*AggregatorRunner
}
