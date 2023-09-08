package middlewares

import (
	"fmt"
	"os"

	"github.com/Cracker-TG/portfolio-service/config"
	"github.com/gin-gonic/gin"
)

func Cors(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("MODE") == "PRODUCTION" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", config.APP_DOMAIN)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding,  Authorization, accept, origin, Cache-Control, cf-turnstile-token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		fmt.Println("OPTIONS-", c.Request.Method)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
