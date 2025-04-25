package logger

import (
	log "log/slog"
	"os"
)

func newSlog(env string) *slog {
	var l *log.Logger

	switch env {
	case "DEV":
		l = log.New(log.NewTextHandler(os.Stdout, &log.HandlerOptions{
			Level: log.LevelDebug,
		}))
	case "PROD":
		panic("unimplemented")
	default:
		panic("invalid env type")
	}

	return &slog{
		log: l,
	}
}

type slog struct {
	log *log.Logger
}

func (s slog) Info(msg string, args ...arg) {
	if len(args) == 0 {
		s.log.Info(msg)
	} else {
		slogArgs := make([]any, 0, len(args))
		for _, arg := range args {
			slogArgs = append(slogArgs, log.Any(arg.key, arg.val))
		}

		s.log.Info(msg, slogArgs...)
	}
}

func (s slog) Error(msg string, args ...arg) {
	if len(args) == 0 {
		s.log.Error(msg)
	} else {
		slogArgs := make([]any, 0, len(args))
		for _, arg := range args {
			slogArgs = append(slogArgs, log.Any(arg.key, arg.val))
		}

		s.log.Error(msg, slogArgs...)
	}
}

func (s slog) Debug(msg string, args ...arg) {
	if len(args) == 0 {
		s.log.Debug(msg)
	} else {
		slogArgs := make([]any, 0, len(args))
		for _, arg := range args {
			slogArgs = append(slogArgs, log.Any(arg.key, arg.val))
		}

		s.log.Debug(msg, slogArgs...)
	}
}

func (s slog) Warn(msg string, args ...arg) {
	if len(args) == 0 {
		s.log.Warn(msg)
	} else {
		slogArgs := make([]any, 0, len(args))
		for _, arg := range args {
			slogArgs = append(slogArgs, log.Any(arg.key, arg.val))
		}

		s.log.Warn(msg, slogArgs...)
	}
}

func (s slog) Close() error {
	//TODO implement me
	panic("implement me")
}
