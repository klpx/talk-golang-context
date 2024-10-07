package metrics

import (
	"context"
	"fmt"
	"github.com/klpx/talk-golang-context/pkg/ctxstore"
	"github.com/redis/go-redis/v9"
	"time"
)

var QueryName = ctxstore.MakeStore[string]()

type RedisMetrics struct {
}

func (r RedisMetrics) DialHook(next redis.DialHook) redis.DialHook {
	return next
}

func (r RedisMetrics) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		start := time.Now()
		err := next(ctx, cmd)
		duration := time.Since(start)
		fmt.Printf("request %s took %s\n", cmd.Name(), duration.String())
		return err
	}
}

func (r RedisMetrics) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return next
}

var _ redis.Hook = &RedisMetrics{}
