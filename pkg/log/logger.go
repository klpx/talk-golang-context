package log

import (
	"context"
	"fmt"
	"log"
)

type Logger struct {
	log     *log.Logger
	extract func(ctx context.Context) map[string]string
}

func MakeLogger(extract func(ctx context.Context) map[string]string) *Logger {
	return &Logger{
		log:     log.Default(),
		extract: extract,
	}
}

func (l *Logger) Infoc(ctx context.Context, msg string, args ...any) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	metaInfo := l.extract(ctx)
	for key, value := range metaInfo {
		msg += fmt.Sprintf("\t%s=%s", key, value)
	}
	l.log.Println(msg)
}
