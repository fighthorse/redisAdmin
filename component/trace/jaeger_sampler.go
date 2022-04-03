package trace

import (
	"math"

	"github.com/uber/jaeger-client-go"
)

// ProbabilisticSampler is a sampler that randomly samples a certain percentage of traces.
type ProbabilisticSampler struct {
	samplingRate     float64
	samplingBoundary uint64
	tags             []jaeger.Tag
}

const maxRandomNumber = 0xffffffff

// NewProbabilisticSampler creates a sampler that randomly samples a certain percentage of traces specified by the
// samplingRate, in the range between 0.0 and 1.0.
//
// It relies on the fact that new trace IDs are 63bit random numbers themselves, thus making the sampling decision
// without generating a new random number, but simply calculating if traceID < (samplingRate * 2^63).
func NewProbabilisticSampler(samplingRate float64) *ProbabilisticSampler {
	if samplingRate < 0.0 || samplingRate > 1.0 {
		samplingRate = 0.01
	}
	return newProbabilisticSampler(samplingRate)
}

func newProbabilisticSampler(samplingRate float64) *ProbabilisticSampler {
	samplingRate = math.Max(0.0, math.Min(samplingRate, 1.0))
	var tags []jaeger.Tag
	return &ProbabilisticSampler{
		samplingRate:     samplingRate,
		samplingBoundary: uint64(float64(maxRandomNumber) * samplingRate),
		tags:             tags,
	}
}

// SamplingRate returns the sampling probability this sampled was constructed with.
func (s *ProbabilisticSampler) SamplingRate() float64 {
	return s.samplingRate
}

// IsSampled implements IsSampled() of Sampler.
func (s *ProbabilisticSampler) IsSampled(id jaeger.TraceID, operation string) (bool, []jaeger.Tag) {
	return s.samplingBoundary >= (id.Low & 0xffffffff), s.tags
}

// Close implements Close() of Sampler.
func (s *ProbabilisticSampler) Close() {
	// nothing to do
}

// Equal implements Equal() of Sampler.
func (s *ProbabilisticSampler) Equal(other jaeger.Sampler) bool {
	if o, ok := other.(*ProbabilisticSampler); ok {
		return s.samplingBoundary == o.samplingBoundary
	}
	return false
}
