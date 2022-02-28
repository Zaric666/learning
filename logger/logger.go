package logger

import "fmt"

type option struct {
	Level string
}

type Option func(*option)

func WithInfoLevel() Option {
	return func(o *option) {
		o.Level = "info"
	}
}

type Logger struct {
	level string
}

func NewLogger(opts ...Option) (Logger, error) {
	opt := &option{
		Level: "info",
	}
	for _, f := range opts {
		f(opt)
	}
	return Logger{
		level: opt.Level,
	}, nil
}

func (l *Logger) info(s string) string {
	return fmt.Sprintf("[%s] %s", l.level, s)
}
