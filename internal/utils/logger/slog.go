package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	log "log/slog"
	"os"
)

func newSlog(env string) *slog {
	var l slog

	switch env {
	case "DEV":
		l.log = log.New(log.NewTextHandler(os.Stdout, &log.HandlerOptions{
			Level: log.LevelDebug,
		}))
	case "PROD":
		f, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}

		l.file = f
		l.chLogsBatch = make(chan batchLog, batchSize)
		l.ctx, l.cancel = context.WithCancel(context.Background())

		l.log = log.New(log.NewJSONHandler(io.MultiWriter(os.Stdout), &log.HandlerOptions{
			Level: log.LevelInfo,
		}))

		go l.writeLogsToFile()
	default:
		panic("invalid env type")
	}

	return &l
}

type slog struct {
	log *log.Logger

	ctx    context.Context
	cancel func()

	file        io.WriteCloser
	chLogsBatch chan batchLog
}

const (
	batchSize = 4
)

type batchLog struct {
	Msg        string `json:"msg"`
	ArgsString string `json:"args"`
	args       []any
}

func (s slog) Info(msg string, args ...arg) {
	if len(args) == 0 {
		s.log.Info(msg)
		if s.file != nil {
			s.chLogsBatch <- batchLog{
				Msg: msg,
			}
		}
	} else {
		slogArgs := make([]any, 0, len(args))
		for _, arg := range args {
			slogArgs = append(slogArgs, log.Any(arg.key, arg.val))
		}

		s.log.Info(msg, slogArgs...)
		if s.file != nil {
			s.chLogsBatch <- batchLog{
				Msg:  msg,
				args: slogArgs,
			}
		}
	}
}

func (s slog) Error(msg string, args ...arg) {
	if len(args) == 0 {
		s.log.Error(msg)
		if s.file != nil {
			s.chLogsBatch <- batchLog{
				Msg: msg,
			}
		}
	} else {
		slogArgs := make([]any, 0, len(args))
		for _, arg := range args {
			slogArgs = append(slogArgs, log.Any(arg.key, arg.val))
		}

		s.log.Error(msg, slogArgs...)
		if s.file != nil {
			s.chLogsBatch <- batchLog{
				Msg:  msg,
				args: slogArgs,
			}
		}
	}
}

func (s slog) Debug(msg string, args ...arg) {
	if len(args) == 0 {
		s.log.Debug(msg)
		if s.file != nil {
			s.chLogsBatch <- batchLog{
				Msg: msg,
			}
		}
	} else {
		slogArgs := make([]any, 0, len(args))
		for _, arg := range args {
			slogArgs = append(slogArgs, log.Any(arg.key, arg.val))
		}

		s.log.Debug(msg, slogArgs...)
		if s.file != nil {
			s.chLogsBatch <- batchLog{
				Msg:  msg,
				args: slogArgs,
			}
		}
	}
}

func (s slog) Warn(msg string, args ...arg) {
	if len(args) == 0 {
		s.log.Warn(msg)
		if s.file != nil {
			s.chLogsBatch <- batchLog{
				Msg: msg,
			}
		}
	} else {
		slogArgs := make([]any, 0, len(args))
		for _, arg := range args {
			slogArgs = append(slogArgs, log.Any(arg.key, arg.val))
		}

		s.log.Warn(msg, slogArgs...)
		if s.file != nil {
			s.chLogsBatch <- batchLog{
				Msg:  msg,
				args: slogArgs,
			}
		}
	}
}

func (s slog) Close() error {
	close(s.chLogsBatch)
	<-s.ctx.Done()
	return nil
}

func (s slog) writeLogsToFile() {
	defer func() {
		_ = s.file.Close()
		s.cancel()
	}()

	batch := make([]batchLog, 0, batchSize)

	for l := range s.chLogsBatch {
		l.ArgsString = fmt.Sprintf("%v", l.args)
		batch = append(batch, l)

		if len(batch) == batchSize {
			b, err := json.Marshal(batch)
			if err != nil {
				batch = batch[:0]
				s.log.Error(fmt.Sprintf("failed to marshal batch: %s", err.Error()))
				continue
			}

			if _, err = s.file.Write(b); err != nil {
				s.log.Error(fmt.Sprintf("failed to write batch: %s", err.Error()))
			}

			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		b, err := json.Marshal(batch)
		if err != nil {
			s.log.Error(fmt.Sprintf("failed to marshal batch: %s", err.Error()))
			return
		}

		if _, err = s.file.Write(b); err != nil {
			s.log.Error(fmt.Sprintf("failed to write batch: %s", err.Error()))
		}
	}
}
