package services

import (
	"github.com/rssed-day/rssed-day-core/context"
	"github.com/rssed-day/rssed-day-core/pipeline"
)

// PipelineService -
type PipelineService struct {
}

// NewPipelineService -
func NewPipelineService() *PipelineService {
	return &PipelineService{}
}

// Pipe -
func (p *PipelineService) Pipe(factory context.ConfigFactory) error {
	cfg, err := factory.Config()
	if err != nil {
		return err
	}

	ctx, err := context.NewContext(cfg)
	if err != nil {
		return err
	}

	return pipeline.NewPipeline(ctx).Pipe(nil)
}
