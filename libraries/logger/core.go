package logger

import (
	"context"
	"fmt"
	lr "github.com/sirupsen/logrus"
	"io"
	"runtime"
)

var (
	ErrUnknownFormat = fmt.Errorf(`[UNKNOWN LOG FORMAT] [FAILED] Logger Error`)
)

type (
	ILogger interface {
		InfoWithContext(ctx context.Context, uuid, msg string, values ...interface{})
		ErrorWithContext(ctx context.Context, uuid string, err error, msg string, values ...interface{})
		DebugWithContext(ctx context.Context, uuid, msg string, values ...interface{})
		Info(uuid, msg string, values ...interface{})
		Error(uuid string, err error, msg string, values ...interface{})
		Debug(uuid, msg string, values ...interface{})
	}
	Logger struct {
		Loggers       lr.Logger
		AppName       string
		DefaultFields map[string]interface{}
		ContextFields map[string]interface{}
		SkipCaller    int
		Formatter     string
	}

	Option struct {
		Level         string
		Out           io.Writer
		AppName       string
		DefaultFields map[string]interface{}
		Formatter     string
	}
)

func NewLogger(opt *Option) ILogger {
	if opt.DefaultFields == nil {
		opt.DefaultFields = map[string]interface{}{}
	}

	log := lr.New()
	switch opt.Formatter {
	case FormatJSON:
		log.SetFormatter(&lr.JSONFormatter{})
	case FormatText:
		log.SetFormatter(&lr.TextFormatter{})
	default:
		log.Panic(ErrUnknownFormat)
	}

	if opt.AppName != "" {
		opt.DefaultFields["service_name"] = opt.AppName
	} else {
		opt.DefaultFields["service_name"] = "unknown"
	}

	return &Logger{
		Loggers:       *log,
		DefaultFields: opt.DefaultFields,
		SkipCaller:    0,
		Formatter:     opt.Formatter,
	}
}

func (l *Logger) collectFields(ctx context.Context, values ...LogField) lr.Fields {

	fields := make(lr.Fields)
	l.setDefaultFields(fields)
	if ctx != nil && ctx != context.TODO() {
		l.setContextFields(ctx, fields)
	}
	l.setAdditionalLogFields(fields, values...)
	l.setCaller(fields)

	return fields
}

func (l *Logger) setDefaultFields(fields lr.Fields) {
	for k, v := range l.DefaultFields {
		fields[k] = v
	}
}

func (l *Logger) setContextFields(ctx context.Context, fields lr.Fields) {
	if ctx != nil {
		for k, v := range l.ContextFields {
			if val := ctx.Value(v); val != nil {
				fields[k] = val
			}
		}
	}
}

func (l *Logger) setAdditionalLogFields(fields lr.Fields, values ...LogField) {
	for _, field := range values {
		v := field.Value
		switch vt := v.(type) {
		case nil, string, int, int32, int64, float64, bool:
			fields[field.Key] = vt
		default:
			fields[field.Key] = fmt.Sprintf("%+v", vt)
		}
	}
}

const layerCountInThisLibrary = 3

func (l *Logger) setCaller(fields lr.Fields) {

	// Disable caller key when opt.SkipCaller < 0
	if l.SkipCaller >= 0 {
		if _, file, line, ok := runtime.Caller(l.SkipCaller + layerCountInThisLibrary); ok {
			fields["caller"] = fmt.Sprintf("%s:%d", file, line)
		}
	}
}
