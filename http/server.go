package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var SrvHandler *http.Server

func Run() {
	gin.SetMode(viper.GetString("global.server_mode"))
	r := InitRouter(InitMiddlewares()...)
	SrvHandler = &http.Server{
		Addr:         viper.GetString("http.addr"),
		Handler:      r,
		ReadTimeout:  time.Duration(viper.GetInt("http.read_timeout")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt("http.write_timeout")) * time.Second,
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", viper.GetString("http.addr"))
		if err := SrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", viper.GetString("http.addr"), err)
		}
	}()
}

func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := SrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop:%s err:%v\n", viper.GetString("http.addr"), err)
	}
	log.Printf(" [INFO] HttpServerStop:%s stopped\n", viper.GetString("http.addr"))
}
