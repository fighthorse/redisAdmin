package log

import (
	"context"
)

// context key

type ctxKey int

const (
	ContextKeyTraceId ctxKey = iota
)

func WithTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, ContextKeyTraceId, traceId)
}

func TraceIdFromCtx(ctx context.Context) (traceId string) {
	traceId, _ = ctx.Value(ContextKeyTraceId).(string)
	return
}
