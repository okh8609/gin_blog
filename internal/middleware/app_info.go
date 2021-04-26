package middleware

import "github.com/gin-gonic/gin"

func AppInfo(c *gin.Context) {
	c.Set("app_name", "blog-service")
	c.Set("app_version", "1.0.0")
	c.Next()
}
