package log

import (
	"context"
)

type Logger interface {
	Debug(ctx context.Context, msg string, keyvals ...interface{})
	Info(ctx context.Context, msg string, keyvals ...interface{})
	Error(ctx context.Context, message string, keyvals ...interface{})
}

func NewLogger() Logger {
	return &logger{
		adapter: newSlogAdapter(),
	}
}

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
)

type adapter interface {
	Log(level Level, msg string, keyvals ...interface{})
}

type logger struct {
	adapter adapter
}

func (l *logger) Debug(_ context.Context, msg string, keyvals ...interface{}) {
	l.adapter.Log(LevelDebug, msg, keyvals)
}

func (l *logger) Info(_ context.Context, msg string, keyvals ...interface{}) {
	l.adapter.Log(LevelInfo, msg, keyvals)
}

func (l *logger) Error(_ context.Context, msg string, keyvals ...interface{}) {
	l.adapter.Log(LevelError, msg, keyvals)
}
