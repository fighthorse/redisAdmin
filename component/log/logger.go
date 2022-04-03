package log

import (
	"github.com/dolab/logger"
)

// A LogLevelType defines the level logging should be performed at. Used to instruct
// the SDK which statements should be logged.
type LogLevelType uint

// LogLevel returns the pointer to a LogLevel. Should be used to workaround
// not being able to take the address of a non-composite literal.
func LogLevel(l LogLevelType) *LogLevelType {
	return &l
}

// Value returns the LogLevel value or the default value LogOff if the LogLevel
// is nil. Safe to use on nil value LogLevelTypes.
func (l *LogLevelType) Value() LogLevelType {
	if l != nil {
		return *l
	}
	return LogOff
}

// Matches returns true if the v LogLevel is enabled by this LogLevel. Should be
// used with logging sub levels. Is safe to use on nil value LogLevelTypes. If
// LogLevel is nil, will default to LogOff comparison.
func (l *LogLevelType) Matches(v LogLevelType) bool {
	c := l.Value()
	return c&v == v
}

// AtLeast returns true if this LogLevel is at least high enough to satisfies v.
// Is safe to use on nil value LogLevelTypes. If LogLevel is nil, will default
// to LogOff comparison.
func (l *LogLevelType) AtLeast(v LogLevelType) bool {
	c := l.Value()
	return c >= v
}

const (
	// LogOff states that no logging should be performed by the SDK. This is the
	// default state of the SDK, and should be use to disable all logging.
	LogOff LogLevelType = iota * 0x1000

	// LogDebug state that debug output should be logged by the SDK. This should
	// be used to inspect request made and responses received.
	LogDebug
)

// log levels
const (
	LogInfo = LogDebug | (1 << iota)
	LogWarn
	LogError
)

// A dummyLogger implements qulibs Logger interface.
type dummyLogger struct {
	level LogLevelType
	log   *logger.Logger
}

// NewLogger returns a Logger with level given.
func NewLogger(level LogLevelType) Logger {
	log, _ := logger.New("stderr")
	log.SetSkip(3)

	return &dummyLogger{
		level: level,
		log:   log,
	}
}

// NewDummyLogger returns a Logger with LogOff level, which means no logs will be wrote.
func NewDummyLogger() Logger {
	return NewLogger(LogOff)
}

// Debug of dummy logger
func (l *dummyLogger) Debug(args ...interface{}) {
	if l.level.AtLeast(LogDebug) {
		return
	}

	l.log.Debug(args...)
}

// Debugf of dummy logger
func (l *dummyLogger) Debugf(format string, args ...interface{}) {
	if l.level.AtLeast(LogDebug) {
		return
	}
	if format[len(format)-1] != '\n' {
		format += "\n"
	}

	l.log.Debugf(format, args...)
}

// Info of dummy logger
func (l *dummyLogger) Info(args ...interface{}) {
	if l.level.AtLeast(LogInfo) {
		return
	}

	l.log.Info(args...)
}

// Infof of dummy logger
func (l *dummyLogger) Infof(format string, args ...interface{}) {
	if l.level.AtLeast(LogInfo) {
		return
	}

	l.log.Infof(format, args...)
}

// Warn of dummy logger
func (l *dummyLogger) Warn(args ...interface{}) {
	if l.level.AtLeast(LogWarn) {
		return
	}

	l.log.Warn(args...)
}

// Warnf of dummy logger
func (l *dummyLogger) Warnf(format string, args ...interface{}) {
	if l.level.AtLeast(LogWarn) {
		return
	}

	l.log.Warnf(format, args...)
}

// Error of dummy logger
func (l *dummyLogger) Error(args ...interface{}) {
	if l.level.AtLeast(LogError) {
		return
	}

	l.log.Error(args...)
}

// Errorf of dummy logger
func (l *dummyLogger) Errorf(format string, args ...interface{}) {
	if l.level.AtLeast(LogError) {
		return
	}

	l.log.Errorf(format, args...)
}
