package errors

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
)

const (
	ErrInternal = iota
	ErrBadRequest
)

type customError struct {
	msg      string `json:"msg"`
	code     uint   `json:"code"`
	location string `json:"location"`
}

func (e *customError) Error() string {
	return e.msg
}

func New(msg string, code uint) error {
	_, file, line, _ := runtime.Caller(1)

	return &customError{
		msg:      msg,
		code:     code,
		location: fmt.Sprintf("%s:%d", file, line),
	}
}

const (
	HTTP = iota
)

func Parse(t uint, err error) (string, string, uint) {
	switch t {
	case HTTP:
		var e *customError
		if !errors.As(err, &e) {
			return e.location, err.Error(), http.StatusInternalServerError
		}

		switch e.code {
		case ErrInternal:
			return e.location, err.Error(), http.StatusInternalServerError
		case ErrBadRequest:
			return "", e.msg, http.StatusBadRequest
		default:
			return e.location, err.Error(), http.StatusInternalServerError
		}
	default:
		panic("unknown error parse type")
	}
}
