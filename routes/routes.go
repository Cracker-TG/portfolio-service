package routes

import (
	"github.com/Cracker-TG/portfolio-service/controllers/backend/user"
	"github.com/Cracker-TG/portfolio-service/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	userController := new(user.UserController)

	v1 := router.Group("api/v1")
	{
		backend := v1.Group("backend")
		{
			backend.POST("/login", userController.Login)
			authorized := backend.Group("/")
			authorized.Use(middlewares.AuthRequired())
			{
				authorized.GET("/user/info", userController.Info)
			}
		}
	}

	return router
}
