package trace

import (
	"fmt"
	"time"

	"github.com/uber/jaeger-client-go"
	j "github.com/uber/jaeger-client-go/thrift-gen/jaeger"
)

type LocalSpan map[string]interface{}

func NewSpan() LocalSpan {
	return make(LocalSpan)
}

func (s LocalSpan) Set(key string, value interface{}) {
	s[key] = value
}

func (s LocalSpan) SetTraceID(traceIDHigh, traceIDLow int64) {
	var traceID string
	if traceIDHigh != 0 {
		traceID = fmt.Sprintf("%016x%016x", uint64(traceIDHigh), uint64(traceIDLow))
	} else {
		traceID = fmt.Sprintf("%016x", uint64(traceIDLow))
	}
	s.Set("traceID", traceID)
}

func (s LocalSpan) SetParentSpanID(parentSpanID int64) {
	if parentSpanID != 0 {
		s.Set("parentSpanID", fmt.Sprintf("%016x", uint64(parentSpanID)))
	} else {
		s.Set("parentSpanID", fmt.Sprintf("%x", uint64(parentSpanID)))
	}
}

func (s LocalSpan) SetSpanID(spanID int64) {
	s.Set("spanID", fmt.Sprintf("%016x", uint64(spanID)))
}

func (s LocalSpan) SetOperationName(operationName string) {
	s.Set("operationName", operationName)
}

func (s LocalSpan) SetTag(key string, value string) {
	s.Set("tags."+key, value)
}

func (s LocalSpan) SetServiceName(serviceName string) {
	s.Set("process.serviceName", serviceName)
}

func (s LocalSpan) SetHostname(hostname string) {
	s.Set("process.tags.hostname", hostname)
}

func (s LocalSpan) SetIP(ip string) {
	s.Set("process.tags.ip", ip)
}

func BuildSpan(span *jaeger.Span) LocalSpan {
	t := jaeger.BuildJaegerThrift(span)
	x := NewSpan()
	x.SetTraceID(t.TraceIdHigh, t.TraceIdLow)
	x.SetParentSpanID(t.ParentSpanId)
	x.SetSpanID(t.SpanId)
	x.Set("startTime", t.StartTime)
	x.Set("duration", t.Duration)
	x.SetOperationName(t.OperationName)

	process := jaeger.BuildJaegerProcessThrift(span)
	x.SetServiceName(process.GetServiceName())

	for _, tag := range process.GetTags() {
		switch tag.Key {
		case jaeger.TracerHostnameTagKey:
			x.SetHostname(tag.GetVStr())
		case jaeger.TracerIPTagKey:
			x.SetIP(tag.GetVStr())
		}
	}

	for _, tag := range t.Tags {
		var value string
		switch {
		case tag.VBinary != nil:
			value = fmt.Sprintf("%s", tag.VBinary)
		case tag.VBool != nil:
			if *tag.VBool {
				value = "1"
			} else {
				value = "0"
			}
		case tag.VDouble != nil:
			value = fmt.Sprintf("%v", *tag.VDouble)
		case tag.VLong != nil:
			value = fmt.Sprintf("%v", *tag.VLong)
		case tag.VStr != nil:
			value = fmt.Sprintf("%s", *tag.VStr)
		}
		x.SetTag(tag.Key, value)
	}

	var logs []*Log
	for _, log := range t.Logs {
		t := time.Unix(log.Timestamp/1000000, log.Timestamp%1000000*1000).UTC()
		fields := make([]*Field, 0, len(log.Fields))
		for _, field := range log.Fields {
			fields = append(fields, buildField(field))
		}
		logs = append(logs, &Log{
			Timestamp: t.Format("2006-01-02T15:04:05.000000Z"),
			Fields:    fields,
		})
	}
	if logs != nil {
		x.Set("logs", logs)
	}

	return x
}

type Log struct {
	Timestamp string   `json:"timestamp"`
	Fields    []*Field `json:"fields"`
}

type Field struct {
	Key     string   `json:"key"`
	VType   string   `json:"vType"`
	VStr    *string  `json:"vStr,omitempty"`
	VDouble *float64 `json:"vDouble,omitempty"`
	VBool   *bool    `json:"vBool,omitempty"`
	VLong   *int64   `json:"vLong,omitempty"`
	VBinary []byte   `json:"vBinary,omitempty"`
}

func buildField(tag *j.Tag) *Field {
	field := &Field{
		Key:     tag.Key,
		VStr:    tag.VStr,
		VDouble: tag.VDouble,
		VBool:   tag.VBool,
		VLong:   tag.VLong,
		VBinary: tag.VBinary,
	}
	switch {
	case tag.VBinary != nil:
		field.VType = "binary"
	case tag.VBool != nil:
		field.VType = "bool"
	case tag.VDouble != nil:
		field.VType = "float64"
	case tag.VLong != nil:
		field.VType = "int64"
	case tag.VStr != nil:
		field.VType = "string"
	}
	return field
}
