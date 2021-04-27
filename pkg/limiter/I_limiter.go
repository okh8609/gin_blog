package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type ILimiter interface {
	// 新增多個 TokenBucket
	AddBuckets(rules ...BucketInfo) ILimiter
	// 取得對應的限流器的名稱
	GetKey(c *gin.Context) string
	// 取得 TokenBucket
	GetBucket(key string) (*ratelimit.Bucket, bool)
}

type Limiter struct {
	buckets map[string]*ratelimit.Bucket
}

type BucketInfo struct {
	Name      string        // 該 TokenBucket 的名稱
	Capacity  int64         // 該 TokenBucket 的總容量
	Frequency time.Duration // 多久放一次 Token
	Quantum   int64         // 一次放入幾個 Token
}
