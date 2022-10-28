package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rssed-day/rssed-day-core/context"
	"github.com/rssed-day/rssed-day-core/http/dtos"
	"github.com/rssed-day/rssed-day-core/services"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	PipelineActionPipe = "pipe"
)

// PipelineAction -
func PipelineAction(ctx *gin.Context) {
	var (
		req dtos.PipelineActionModel
		err error
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		logrus.Errorf(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	switch req.Action {
	case PipelineActionPipe:
		if err := services.NewPipelineService().Pipe(context.NewBaseConfigFactory(&req.Config)); err != nil {
			logrus.Errorf(err.Error())
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	default:
		err := errors.New(fmt.Sprintf("pipeline action %s not support", req.Action))
		logrus.Errorf(err.Error())
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, struct{}{})
}
