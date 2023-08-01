package routes

import (
	"net/http"

	"github.com/Cracker-TG/portfolio-service/controllers/contacts"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	contactController := new(contacts.ContactController)

	v1 := router.Group("api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world")
		})
		users := v1.Group("api/v1")
		{
			users.POST("/contacts", contactController.Store)
		}
	}

	return router
}
