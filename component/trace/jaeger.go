package trace

import (
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/utils"
	"github.com/uber/jaeger-client-go/zipkin"
)

var rng *rand.Rand

func init() {
	rng = utils.NewRand(time.Now().UnixNano() ^ int64(ip()))
}

// NewJaegerTracer构造返回一个opentracing.Tracer。采样率小于0会被修正为0; 修正后的采样率
// 大于1，使用限速采样，每秒最多采样指定数量的trace; 修正后的采样率介于0和1之间则使用概率采样。
func NewJaegerTracer(serviceName, traceFileName string, samplingRate float64, sig os.Signal, bufferSize int, options ...jaeger.TracerOption) (opentracing.Tracer, io.Closer, error) {
	var sampler jaeger.Sampler
	if samplingRate < 0 {
		samplingRate = 0
	}
	if samplingRate <= 1 {
		sampler = NewProbabilisticSampler(samplingRate)
	} else {
		sampler = jaeger.NewRateLimitingSampler(samplingRate)
	}

	reporter := NewJaegerFileReporter(traceFileName, JaegerReopenSignal(sig), JaegerBufferSize(bufferSize))

	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()

	injector := jaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator)
	extractor := jaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator)

	throtter := jaeger.TracerOptions.DebugThrottler(DefaultThrottler{})

	randomNumber := jaeger.TracerOptions.RandomNumber(defaultRandomNumber)

	options = append(options, injector, extractor, throtter, randomNumber)

	// create Jaeger tracer
	tracer, closer := jaeger.NewTracer(
		serviceName,
		sampler,
		reporter,
		options...,
	)

	//opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}

// randomNumber 冲突解决随机算法
func defaultRandomNumber() uint64 {
	// 0x1000000000000000 保证转化成16进制没有前缀0
	return uint64(rng.Int63() | 0x1000000000000000)
}

func ip() uint32 {
	ip, err := utils.HostIP()
	if err != nil {
		return 0
	}
	value := utils.PackIPAsUint32(ip)
	return value
}

// DefaultThrottler doesn't throttle at all.
type DefaultThrottler struct{}

// IsAllowed implements Throttler#IsAllowed.
func (t DefaultThrottler) IsAllowed(operation string) bool {
	return true
}

func buildSampler(samplingRate float64) jaeger.Sampler {
	if samplingRate > 1.0 {
		return jaeger.NewRateLimitingSampler(samplingRate)
	}
	if samplingRate < 0 {
		samplingRate = 0
	}
	sampler, _ := jaeger.NewProbabilisticSampler(samplingRate)
	return sampler
}
