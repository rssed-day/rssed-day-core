package aggregators

import "github.com/rssed-day/rssed-day-core/models"

// AggregatorActions -
//                 ┌────────────┐
//            ┌──▶ │ Aggregator │───┐
//            │    └────────────┘   │
//  ______    │    ┌────────────┐   │     ______
// ()_____)───┼──▶ │ Aggregator │───┼──▶ ()_____)
//            │    └────────────┘   │
//            │    ┌────────────┐   │
//            └──▶ │ Aggregator │───┘
//                 └────────────┘
type AggregatorActions struct {
	Src         <-chan models.Object
	Dst         chan<- models.Object
	Aggregators []*AggregatorRunner
}
