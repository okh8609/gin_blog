package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
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

	var traceID string
	var spanID string
	var jaegerSpanContext = span.Context()
	switch jaegerSpanContext.(type) {
	case jaeger.SpanContext:
		jaegerContext:=jaegerSpanContext.(jaeger.SpanContext)
		traceID = jaegerContext.TraceID().String()
		spanID = jaegerContext.SpanID().String()
	}

	c.Set("X-Trace-ID",traceID)
	c.Set("X-Span-ID",spanID)
	c.Request = c.Request.WithContext(newContext)
	c.Next()
}
