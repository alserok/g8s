package logger

type Logger interface {
	Info(msg string, args ...arg)
	Error(msg string, args ...arg)
	Debug(msg string, args ...arg)
	Warn(msg string, args ...arg)

	Close() error
}

const (
	Slog = iota
)

func New(t uint, env string) Logger {
	switch t {
	case Slog:
		return newSlog(env)
	default:
		panic("invalid logger type")
	}
}

func WithArg(key string, val any) arg {
	return arg{
		key: key,
		val: val,
	}
}

type arg struct {
	key string
	val any
}
