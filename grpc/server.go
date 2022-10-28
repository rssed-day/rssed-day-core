package grpc

import (
	"context"
	"fmt"
	pb "github.com/rssed-day/rssed-day-core/grpc/protos/pipeline"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var SrvHandler *grpc.Server

func Run() {
	SrvHandler = grpc.NewServer()
	pb.RegisterPipelineServiceServer(SrvHandler, &pb.PipelineHandler{})
	go func() {
		log.Printf(" [INFO] GrpcServerRun:%s\n", viper.GetString("grpc.addr"))
		lis, err := net.Listen("tcp", fmt.Sprintf("%s", viper.GetString("grpc.addr")))
		if err != nil {
			log.Fatalf("[ERROR] GrpcServerRun:%s err:%v\n", viper.GetString("grpc.addr"), err)
		}
		if err = SrvHandler.Serve(lis); err != nil {
			log.Fatalf(" [ERROR] GrpcServerRun:%s err:%v\n", viper.GetString("grpc.addr"), err)
		}
	}()
}

func Stop() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	SrvHandler.Stop()
	log.Printf(" [INFO] GrpcServerStop:%s stopped\n", viper.GetString("grpc.addr"))
}
