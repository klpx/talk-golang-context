package storage

import (
	"context"
	"errors"
	"github.com/klpx/talk-golang-context/pkg/metrics"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	rediscl *redis.Client
}

func Make(rc *redis.Client) *Storage {
	return &Storage{
		rediscl: rc,
	}
}

func (s *Storage) RecordVisit(ctx context.Context, name string) (int64, int64, error) {
	userCount, err1 := s.rediscl.Incr(
		metrics.QueryName.WithValue(ctx, "user_inc"),
		"visitor."+name,
	).Result()
	totalCount, err2 := s.rediscl.Incr(
		metrics.QueryName.WithValue(ctx, "total_inc"),
		"total_visits",
	).Result()
	return userCount, totalCount, errors.Join(err1, err2)
}
