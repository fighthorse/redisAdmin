package log

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"

	"github.com/fighthorse/redisAdmin/pkg/conf"
)

type (
	srvlogger struct {
		//AppLogger 业务日志句柄（非实时， 100ms刷新一次）
		appLogger zerolog.Logger
		//AccessLogger 访问日志句柄（非实时， 100ms刷新一次）
		accessLogger zerolog.Logger
	}
)

var (
	SrvLogger = &srvlogger{
		appLogger: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
)

func Init() {
	{
		// 初始化业务日志
		fileName := conf.GConfig.Log.App.FilePath
		fmt.Println("AppLogPath:" + fileName)
		err := os.MkdirAll(path.Dir(fileName), 0777)
		if err != nil {
			panic(fmt.Errorf("creat log dir error: %s", err))
		}
		f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Errorf("open log file error: %s", err))
		}
		zerolog.TimeFieldFormat = time.RFC3339

		// 1000000 * 1024 /1024/1024 ~= 988M [5m]
		w := diode.NewWriter(f, 1000000, 100*time.Millisecond, func(missed int) {
			SrvLogger.appLogger.Log().Int("count", missed).Msg("app_log_miss")
		})

		// level
		l := parseLevel(conf.GConfig.Log.App.Level)
		SrvLogger.appLogger = zerolog.New(w).Level(l).With().Timestamp().Logger()
	}

	{
		// 初始化访问日志
		fileName := conf.GConfig.Log.Access.FilePath
		fmt.Println("AppAccessLogPath:" + fileName)
		err := os.MkdirAll(path.Dir(fileName), 0777)
		if err != nil {
			panic(fmt.Errorf("creat log dir error: %s", err))
		}
		f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Errorf("open log file error: %s", err))
		}
		w := diode.NewWriter(f, 1000000, 100*time.Millisecond, func(missed int) {
			SrvLogger.accessLogger.Log().Int("count", missed).Msg("access_log_miss")
		})
		SrvLogger.accessLogger = zerolog.New(w).With().Timestamp().Logger()

	}
}

type Fields map[string]interface{}

func (f Fields) putInContext() map[string]interface{} {
	if len(f) == 0 {
		return f
	}
	m := make(map[string]interface{}, len(f))
	m["context"] = f
	return m
}

func Debug(ctx context.Context, msg string, f Fields) {
	f["caller"] = caller(1)
	withTraceIdLogger(ctx).Debug().Fields(f.putInContext()).Msg(msg)
}

func Info(ctx context.Context, msg string, f Fields) {
	f["caller"] = caller(1)
	withTraceIdLogger(ctx).Info().Fields(f.putInContext()).Msg(msg)
}

func Warn(ctx context.Context, msg string, f Fields) {
	f["caller"] = caller(1)
	withTraceIdLogger(ctx).Warn().Fields(f.putInContext()).Msg(msg)
}

func Error(ctx context.Context, msg string, f Fields) {
	f["caller"] = caller(1)
	withTraceIdLogger(ctx).Error().Fields(f.putInContext()).Msg(msg)
}

func Fatal(ctx context.Context, msg string, f Fields) {
	f["caller"] = caller(1)
	withTraceIdLogger(ctx).Fatal().Fields(f.putInContext()).Msg(msg)
}

// Panic 期望的panic
type expectedPanic struct{}

var ExpectedPanic = expectedPanic{}

func Panic(ctx context.Context, msg string, f Fields) {
	f["caller"] = caller(1)
	withTraceIdLogger(ctx).Panic().Fields(f.putInContext()).Msg(msg)
	panic(ExpectedPanic)
}

func Stack(ctx context.Context, msg string, f Fields) {
	f["stacktrace"] = string(debug.Stack())
	withTraceIdLogger(ctx).Info().Fields(f.putInContext()).Msg(msg)
}

func withTraceIdLogger(ctx context.Context) *zerolog.Logger {
	traceId := TraceIdFromCtx(ctx)
	l := SrvLogger.appLogger.With().Str("trace_id", traceId).Logger()
	return &l
}

func AccessLog(ctx context.Context, fields Fields) {
	traceId := TraceIdFromCtx(ctx)
	l := SrvLogger.accessLogger.With().Str("trace_id", traceId).Fields(fields.putInContext()).Logger()
	l.Log().Msg("")

	//l := SrvLogger.accessLogger.With().Str("trace_id", traceId)
	//for k,v := range fields {
	//	switch v.(type) {
	//	case nil:
	//		l = l.Str(k,"nil")
	//	case int,int32,int64:
	//		l = l.Int(k,v.(int))
	//	case float32,float64:
	//		l = l.Float64(k,v.(float64))
	//	case string:
	//		l = l.Str(k,v.(string))
	//	case bool :
	//		l = l.Bool(k,v.(bool))
	//	default:
	//		continue
	//	}
	//}
	//l2 := l.Logger()
	//l2.Log().Msg("")
}

func AppLog(ctx context.Context) *zerolog.Logger {
	traceId := TraceIdFromCtx(ctx)
	l := SrvLogger.appLogger.With().Str("trace_id", traceId).Logger()
	return &l
}

func parseLevel(l string) (level zerolog.Level) {
	switch strings.ToUpper(l) {
	case "DEBUG":
		level = zerolog.DebugLevel
	case "INFO":
		level = zerolog.InfoLevel
	case "WARN", "WARNING":
		level = zerolog.WarnLevel
	case "ERROR":
		level = zerolog.ErrorLevel
	case "FATAL":
		level = zerolog.FatalLevel
	case "PANIC":
		level = zerolog.PanicLevel
	case "NIL", "NULL", "DISCARD", "NO":
		level = zerolog.Disabled
	default:
		level = zerolog.InfoLevel
	}

	return
}

// caller的显示形式为 File:Line
func caller(depth int) string {
	_, f, n, ok := runtime.Caller(1 + depth)
	if !ok {
		return ""
	}
	if ok {
		idx := strings.LastIndex(f, "github.com")
		if idx >= 0 {
			f = f[idx+10:]
		}
	}
	return fmt.Sprintf("%s:%d", f, n)
}
