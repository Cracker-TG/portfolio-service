package routes

import (
	"net/http"

	"github.com/Cracker-TG/crboard/controllers/users"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userController := new(users.UserController)

	v1 := router.Group("api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world")
		})
		users := v1.Group("users")
		{
			//users.POST("/login", userController.Login)
			users.GET("/info", userController.Info)
			// authorized := backend.Group("/")
			// authorized.Use(middlewares.AuthRequired())
			// {
			// authorized.GET("/user/info", userController.Info)
			// }
		}
	}

	return router
}
