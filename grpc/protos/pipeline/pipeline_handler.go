package pipeline

import (
	"context"
	"errors"
	"fmt"
	context2 "github.com/rssed-day/rssed-day-core/context"
	"github.com/rssed-day/rssed-day-core/services"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	PipelineActionPipe = "pipe"
)

type PipelineHandler struct{}

// mustEmbedUnimplementedPipelineServiceServer -
func (h *PipelineHandler) mustEmbedUnimplementedPipelineServiceServer() {}

// PipelineAction -
func (h *PipelineHandler) PipelineAction(ctx context.Context, in *PipelineActionModel) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	switch in.Action {
	case PipelineActionPipe:
		if err := services.NewPipelineService().Pipe(context2.NewProtoConfigFactory(in.Config)); err != nil {
			logrus.Errorf(err.Error())
			return out, err
		}
	default:
		err := errors.New(fmt.Sprintf("pipeline action %s not support", in.Action))
		logrus.Errorf(err.Error())
		return out, err
	}
	return out, nil
}
