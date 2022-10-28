package test

import (
	"context"
	pb "github.com/rssed-day/rssed-day-core/grpc/protos/pipeline"
	"github.com/rssed-day/rssed-day-core/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
)

func TestGrpcPipe(t *testing.T) {
	conn, err := grpc.Dial("localhost:51402", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	defer conn.Close()

	path, err := utils.Interface2Anypb("assets/helloworld.json")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	client := pb.NewPipelineServiceClient(conn)
	_, err = client.PipelineAction(context.Background(), &pb.PipelineActionModel{
		Action: "pipe",
		Config: &pb.Config{
			Inputs: []*pb.InputConfig{
				{
					Name: "file",
					Args: map[string]*anypb.Any{
						"path": path,
					},
				},
			},
			Outputs: []*pb.OutputConfig{
				{
					Name: "cmd",
				},
			},
		},
	})
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	return
}
