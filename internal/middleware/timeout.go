package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ccc, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()  // releases resources if slowOperation completes before timeout elapses
		c.Request = c.Request.WithContext(ccc)
		c.Next()
	}
}
