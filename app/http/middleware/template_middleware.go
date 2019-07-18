package middleware

import (
	"github.com/gin-gonic/gin"
)


func TemplateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}