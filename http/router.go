package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rssed-day/rssed-day-core/http/controllers"
	"github.com/spf13/viper"
)

const (
	ApiV1Prefix = "/api/v1"
)

func InitMiddlewares() []gin.HandlerFunc {
	var middlewareHandlerFuncs []gin.HandlerFunc
	if viper.GetBool("http.enable_basic_auth") {
		middlewareHandlerFuncs = append(middlewareHandlerFuncs, gin.BasicAuth(gin.Accounts{
			viper.GetString("http.basic_auth_username"): viper.GetString("http.basic_auth_password"),
		}))
	}
	return middlewareHandlerFuncs
}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	// health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//// swagger api routes
	//docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	//docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	//docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ApiV1Groups(router)

	return router
}

// ApiV1Groups -
func ApiV1Groups(router *gin.Engine) *gin.Engine {
	// pipeline api routes
	pipeGroup := router.Group(ApiV1Prefix + "/pipelines")
	{
		// POST /pipeline/action
		pipeGroup.POST("/action", controllers.PipelineAction)
	}
	return router
}
