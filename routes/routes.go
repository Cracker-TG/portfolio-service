package routes

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//router.GET("/ping", main.Status)

	v1 := router.Group("api/v1")
	{
		mainGroup := v1.Group("main")
		{
			mainGroup.GET("/push-noti", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
		}
	}

	return router
}
