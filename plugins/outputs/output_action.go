package outputs

import "github.com/rssed-day/rssed-day-core/models"

// OutputActions -
//
//                            ┌────────┐
//                       ┌──▶ │ OutputPlugin │
//                       │    └────────┘
//  ______     ┌─────┐   │    ┌────────┐
// ()_____)──▶ │ Buf │───┼──▶ │ OutputPlugin │
//             └─────┘   │    └────────┘
//                       │    ┌────────┐
//                       └──▶ │ OutputPlugin │
//                            └────────┘
type OutputActions struct {
	Src     <-chan models.Object
	Outputs []*OutputRunner
}
