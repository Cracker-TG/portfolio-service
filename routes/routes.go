package routes

import (
	"net/http"

	"github.com/Cracker-TG/portfolio-service/controllers/contacts"
	"github.com/Cracker-TG/portfolio-service/validations"
	"github.com/gin-gonic/gin"
)

// ValidationMiddleware is a middleware function to set the custom validator for Gin
func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the custom validator for the context
		c.Set("validator", validations.GetCustomValidator())

		// Continue to the next middleware or handler
		c.Next()
	}
}

func NewRouter(router *gin.Engine) *gin.Engine {
	// router := gin.New()
	// router.Use(CORS())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	contactController := new(contacts.ContactController)

	v1 := router.Group("api/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world")
		})
		contacts := v1.Group("contacts")
		{
			// Create a new validator instance

			// Set the custom validator as the default validator for Gin
			contacts.Use(ValidationMiddleware())

			contacts.POST("", contactController.Store)
		}
	}

	return router
}
