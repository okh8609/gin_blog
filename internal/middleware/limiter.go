package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/pkg/app"
	"github.com/okh8609/gin_blog/pkg/errcode"
	"github.com/okh8609/gin_blog/pkg/limiter"
)

func RateLimiter(l limiter.ILimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.GetKey(c)
		if bucket, ok := l.GetBucket(key); ok { // 找不到Bucket就直接pass(不限流了)
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewGResponse(c)
				response.SendErrResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
