package log

import (
	"log/slog"
	"os"
)

type slogAdapter struct {
	l *slog.Logger
}

func newSlogAdapter() *slogAdapter {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return &slogAdapter{
		l: l,
	}
}

func (s *slogAdapter) Log(level Level, msg string, keyvals ...interface{}) {
	switch level {
	case LevelDebug:
		s.l.Debug(msg, keyvals...)
	case LevelInfo:
		s.l.Info(msg, keyvals...)
	case LevelError:
		s.l.Error(msg, keyvals...)
	default:
	}
}

func (s *slogAdapter) With(keyvals ...interface{}) adapter {
	return &slogAdapter{
		l: s.l.With(keyvals...),
	}
}
