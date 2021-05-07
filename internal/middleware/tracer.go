package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Tracing(c *gin.Context) {
	var newContext context.Context
	var span opentracing.Span

	spanContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(c.Request.Header),
	)

	if err == nil {
		span, newContext = opentracing.StartSpanFromContextWithTracer(
			c.Request.Context(),
			*global.Tracer,
			c.Request.URL.Path,
			opentracing.ChildOf(spanContext),
			opentracing.Tag{
				Key:   string(ext.Component),
				Value: "HTTP",
			},
		)
	} else {
		span, newContext = opentracing.StartSpanFromContextWithTracer(
			c.Request.Context(),
			*global.Tracer,
			c.Request.URL.Path,
		)
	}

	defer span.Finish()
	c.Request = c.Request.WithContext(newContext)
	c.Next()
}
