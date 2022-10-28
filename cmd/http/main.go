package main

import (
	"flag"
	"github.com/rssed-day/rssed-day-core/http"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

var (
	config = flag.String("config", "configs/dev/base.toml", "http config file path")
)

func main() {
	flag.Parse()
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}

	InitHttpServer(*config)
}

func InitHttpServer(config string) {
	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	http.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	http.Stop()
}
