package logger

import (
	"context"
)

type (
	LogField struct {
		Key   string
		Value interface{}
	}
)

func (l *Logger) InfoWithContext(ctx context.Context, uuid, msg string, values ...interface{}) {
	field := l.collectFields(ctx, l.getFields(uuid, values...)...)
	l.Loggers.WithContext(ctx).WithFields(field).Info(msg)
}

func (l *Logger) ErrorWithContext(ctx context.Context, uuid string, err error, msg string, values ...interface{}) {
	field := l.collectFields(ctx, l.getFields(uuid, values...)...)
	l.Loggers.WithContext(ctx).WithFields(field).WithError(err).Error(msg)
}

func (l *Logger) DebugWithContext(ctx context.Context, uuid, msg string, values ...interface{}) {
	field := l.collectFields(ctx, l.getFields(uuid, values...)...)
	l.Loggers.WithContext(ctx).WithFields(field).Debug(msg)
}

func (l *Logger) Info(uuid, msg string, values ...interface{}) {
	field := l.collectFields(context.TODO(), l.getFields(uuid, values...)...)
	l.Loggers.WithFields(field).Info(msg)
}

func (l *Logger) Error(uuid string, err error, msg string, values ...interface{}) {
	field := l.collectFields(context.TODO(), l.getFields(uuid, values...)...)
	l.Loggers.WithFields(field).WithError(err).Error(msg)
}

func (l *Logger) Debug(uuid, msg string, values ...interface{}) {
	field := l.collectFields(context.TODO(), l.getFields(uuid, values...)...)
	l.Loggers.WithFields(field).Debug(msg)
}
