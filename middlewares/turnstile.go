package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func CloudflareTokenValidationMiddleware(secretKey *string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const cloudflareAPIURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"

		token := c.GetHeader("cf-turnstile-token")
		ip := c.ClientIP()
		fmt.Println("token", token, ip)

		formData := url.Values{
			"secret":   {*secretKey},
			"response": {token},
			"remoteip": {ip},
		}

		res, err := http.PostForm(cloudflareAPIURL, formData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
			c.Abort()
			return
		}

		defer res.Body.Close()

		var outcome map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&outcome)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Cloudflare response"})
			c.Abort()
			return
		}

		if success, ok := outcome["success"].(bool); ok && success {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
