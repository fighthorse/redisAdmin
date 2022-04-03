package trace

import (
	"bufio"
	"encoding/json"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/uber/jaeger-client-go"
)

var DefaultSignal os.Signal = syscall.SIGUSR1

// jaegerFileReporter will send spans to a file.
type jaegerFileReporter struct {
	name         string
	file         *os.File
	w            *bufio.Writer
	dataC        chan []byte
	quit         chan struct{}
	signalC      chan os.Signal
	reopenSignal os.Signal
	bufferSize   int
}

// Report implements Report() method of Reporter.
func (r *jaegerFileReporter) Report(span *jaeger.Span) {
	data, err := json.Marshal(BuildSpan(span))
	if err == nil {
		select {
		case r.dataC <- data:
		default:
		}
	}
}

// Close implements Close() method of Reporter.
func (r *jaegerFileReporter) Close() {
	close(r.quit)
	if r.bufferSize > 0 {
		r.w.Flush()
	}
	r.file.Sync()
	r.file.Close()
}

func (r *jaegerFileReporter) writeToFile() {
	for {
		select {
		case data := <-r.dataC:
			data = append(data, '\n')
			r.file.Write(data)
		case <-r.quit:
			return
		case sig := <-r.signalC:
			if sig == r.reopenSignal {
				r.reopen()
			}
		}
	}
}

func (r *jaegerFileReporter) writeToFileWithBuffer() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case data := <-r.dataC:
			data = append(data, '\n')
			r.w.Write(data)
		case <-ticker.C:
			r.w.Flush()
		case <-r.quit:
			return
		case sig := <-r.signalC:
			if sig == r.reopenSignal {
				r.reopen()
			}
		}
	}
}

func (r *jaegerFileReporter) reopen() {
	r.file.Sync()
	r.file.Close()
	file, _ := os.OpenFile(r.name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	r.file = file
	if r.bufferSize > 0 {
		r.w.Flush()
		r.w = bufio.NewWriterSize(r.file, r.bufferSize)
	}
}

// JaegerReporterOption sets a parameter for the HTTP Reporter
type JaegerReporterOption func(r *jaegerFileReporter)

func JaegerReopenSignal(sig os.Signal) JaegerReporterOption {
	return func(r *jaegerFileReporter) { r.reopenSignal = sig }
}

func JaegerBufferSize(size int) JaegerReporterOption {
	return func(r *jaegerFileReporter) { r.bufferSize = size }
}

// NewFileReporter returns a new file Reporter.
func NewJaegerFileReporter(name string, opts ...JaegerReporterOption) jaeger.Reporter {
	os.MkdirAll(filepath.Dir(name), 0755)

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return jaeger.NewNullReporter()
	}

	r := &jaegerFileReporter{
		name:    name,
		file:    file,
		dataC:   make(chan []byte, 65536),
		quit:    make(chan struct{}, 1),
		signalC: make(chan os.Signal),
	}

	for _, opt := range opts {
		opt(r)
	}

	if r.reopenSignal == nil {
		r.reopenSignal = DefaultSignal
	}
	signal.Notify(r.signalC, r.reopenSignal)

	if r.bufferSize > 0 {
		r.w = bufio.NewWriterSize(r.file, r.bufferSize)
		go r.writeToFileWithBuffer()
	} else {
		go r.writeToFile()
	}

	return r
}
