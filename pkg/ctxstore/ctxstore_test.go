package ctxstore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCtxstore(t *testing.T) {
	StoreInt1 := MakeStore[int]()
	StoreInt2 := MakeStore[int]()

	ctx := context.TODO()
	val, exists := StoreInt1.Value(ctx)
	assert.Equal(t, val, 0)
	assert.False(t, exists)

	val, exists = StoreInt1.Value(StoreInt1.WithValue(ctx, 10))
	assert.Equal(t, val, 10)
	assert.True(t, exists)

	// original ctx is untouched
	val, exists = StoreInt1.Value(ctx)
	assert.Equal(t, val, 0)
	assert.False(t, exists)

	ctxDouble := StoreInt2.WithValue(StoreInt1.WithValue(ctx, 100), 200)
	val, exists = StoreInt1.Value(ctxDouble)
	assert.Equal(t, val, 100)
	assert.True(t, exists)
	val, exists = StoreInt2.Value(ctxDouble)
	assert.Equal(t, val, 200)
	assert.True(t, exists)
}
