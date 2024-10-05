package ctxlog

import (
	"context"
	"github.com/klpx/talk-golang-context/pkg/ctxstore"
	"github.com/klpx/talk-golang-context/pkg/log"
)

var globalLogger *log.Logger

func SetGlobalLogger(l *log.Logger) {
	globalLogger = l
}

var loggerStore = ctxstore.MakeStore[*log.Logger]()

func GetLogger(ctx context.Context) *log.Logger {
	if logger, ok := loggerStore.Value(ctx); ok {
		return logger
	}
	return globalLogger
}

func Infof(ctx context.Context, msg string, args ...any) {
	GetLogger(ctx).Infoc(ctx, msg, args...)
}
