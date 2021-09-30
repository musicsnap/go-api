package routers

import (
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-api/controllers"
	_ "go-api/middlewares"
	_ "go-api/middlewares"
	_ "time"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		// test
		testGroup := v1.Group("test")
		{
			test := new(controllers.TestController)
			testGroup.GET("/s", test.Status)
		}
	}
	return router

}
