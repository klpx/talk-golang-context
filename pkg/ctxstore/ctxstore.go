package ctxstore

import (
	"context"
)

func MakeStore[T any]() *ContextStore[T] {
	return &ContextStore[T]{}
}

type ContextStore[T any] struct {
	unique any
}

func (r *ContextStore[T]) Value(ctx context.Context) (T, bool) {
	val, ok := ctx.Value(r).(T)
	return val, ok
}

func (r *ContextStore[T]) WithValue(ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, r, value)
}
