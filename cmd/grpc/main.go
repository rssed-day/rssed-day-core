package main

import (
	"flag"
	"github.com/rssed-day/rssed-day-core/grpc"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

var (
	config = flag.String("config", "configs/dev/base.toml", "grpc config file path")
)

func main() {
	flag.Parse()
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}

	InitGrpcServer(*config)
}

func InitGrpcServer(config string) {
	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	grpc.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	grpc.Stop()
}
