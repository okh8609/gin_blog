package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type MyLimiter struct {
	*Limiter
}

func NewMyLimiter() ILimiter {
	return MyLimiter{
		Limiter: &Limiter{buckets: make(map[string]*ratelimit.Bucket)},
	}
}

func (l MyLimiter) AddBuckets(rules ...BucketInfo) ILimiter {
	for _, rule := range rules {
		if _, ok := l.buckets[rule.Name]; !ok {
			l.buckets[rule.Name] = ratelimit.NewBucketWithQuantum(rule.Frequency, rule.Capacity, rule.Quantum)
		}
	}

	return l
}

func (l MyLimiter) GetKey(c *gin.Context) string {
	// 拿 RequestURI 作為限流的對象，一個URI一個限流器

	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index] // 去掉GET的參數
}

func (l MyLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.buckets[key]
	return bucket, ok
}
