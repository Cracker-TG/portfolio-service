package controllers

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, payload map[string]interface{}) {
	c.JSON(200, gin.H{"success": true, "data": payload})
}

func ErrResponseWithCode(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"success": false, "msg": msg})
}
