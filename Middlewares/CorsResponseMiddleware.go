package Middlewares

import "github.com/gin-gonic/gin"

// CorsMiddleware ...
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// after request
		c.Header("Access-Control-Allow-Origin", "*")
	}
}
