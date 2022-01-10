package middlewares

import (
	"github.com/Cracker-TG/portfolio-service/controllers"
	"github.com/Cracker-TG/portfolio-service/securityTokens"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paseto_maker securityTokens.PasetoInteface = new(securityTokens.PasetoMaker)
		access_token := c.Request.Header.Get("access_token")
		verfify, payload := paseto_maker.VerfifyToken(&access_token)
		if !verfify {
			controllers.ErrResponseWithCode(c, 401, "unauthorized")
			return
		} else {
			c.Set("username", payload.Username)
			c.Set("ID", payload.ID)
			c.Next()
		}
	}
}
