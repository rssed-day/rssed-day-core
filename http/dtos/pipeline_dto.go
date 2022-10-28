package dtos

import "github.com/rssed-day/rssed-day-core/context"

// PipelineActionModel -
type PipelineActionModel struct {
	Action string         `json:"action"`
	Config context.Config `json:"config"`
}
