package srvstatus

import (
	"context"
	"time"
)

func CheckServers(ctx context.Context) []string {
	var result []string
	for i := range 10 {
		err := checkServer(ctx, i)
		if err != nil {
			result = append(result, err.Error())
		} else {
			result = append(result, "ok")
		}
	}
	return result
}

func checkServer(ctx context.Context, i int) error {
	naskolkoTupit := 100 * time.Millisecond
	if i == 2 {
		naskolkoTupit = 500 * time.Millisecond
	}
	var err error

	select {
	case <-ctx.Done():
		err = ctx.Err()
	case <-time.After(naskolkoTupit):

	}
	return err
}
