package trace

import (
	"io"

	"github.com/fighthorse/redisAdmin/component/conf"
	"github.com/opentracing/opentracing-go"
)

var (
	traceServiceName  = "app"
	traceFileName     = "/data/logs/trace/trace_redis.log"
	tracesamplingRate = 0.0001
	traceCloser       io.Closer
)

func Init() {

	traceServiceName = conf.GConfig.Trace.ServiceName
	traceFileName = conf.GConfig.Trace.FilePath
	tracesamplingRate = conf.GConfig.Trace.Sampling

	setGlobalTrace()
}

func setGlobalTrace() {
	if traceCloser != nil {
		traceCloser.Close()
		traceCloser = nil
	}
	tracer, closer, _ := NewJaegerTracer(traceServiceName, traceFileName, tracesamplingRate, nil, 0)
	opentracing.SetGlobalTracer(tracer)
	traceCloser = closer
}
