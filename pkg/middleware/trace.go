package middleware

import (
	"github.com/fighthorse/redisAdmin/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Trace(c *gin.Context) {
	// 每个请求到来时增加trace监控
	ctx := c.Request.Context()
	r := c.Request
	url := r.RequestURI
	if FilterUrl(url) {
		return
	}
	tracer := opentracing.GlobalTracer()
	parentCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	var span opentracing.Span
	if err != nil {
		span = tracer.StartSpan(r.URL.Path)
	} else {
		span = tracer.StartSpan(r.URL.Path, opentracing.ChildOf(parentCtx))
	}

	// set request ID for context
	spanCtx, ok := span.Context().(jaeger.SpanContext)
	if ok {
		ctx = log.WithTraceId(ctx, spanCtx.TraceID().String())
	}

	ext.HTTPMethod.Set(span, r.Method)
	ext.HTTPUrl.Set(span, r.URL.String())
	ext.SpanKindRPCServer.Set(span)
	ctx = opentracing.ContextWithSpan(ctx, span)

	defer span.Finish()

	// reload
	c.Request = r.WithContext(ctx)
	c.Next()

}
