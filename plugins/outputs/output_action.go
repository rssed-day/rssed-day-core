package outputs

import "github.com/rssed-day/rssed-day-core/models"

// OutputActions -
//
//                            ┌────────┐
//                       ┌──▶ │ Output │
//                       │    └────────┘
//  ______     ┌─────┐   │    ┌────────┐
// ()_____)──▶ │ Buf │───┼──▶ │ Output │
//             └─────┘   │    └────────┘
//                       │    ┌────────┐
//                       └──▶ │ Output │
//                            └────────┘
type OutputActions struct {
	Src     <-chan models.Object
	Outputs []*OutputRunner
}
