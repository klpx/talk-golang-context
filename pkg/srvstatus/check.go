package srvstatus

import (
	"context"
	"github.com/klpx/talk-golang-context/pkg/auth"
	"time"
)

func CheckServers(ctx context.Context) ([]string, error) {
	principal, _ := auth.Context.Value(ctx)
	if !principal.IsAuthorized() {
		return nil, auth.ErrNotAuthorized
	}
	var result []string
	for i := range 10 {
		err := checkServer(ctx, i)
		if err != nil {
			result = append(result, err.Error())
		} else {
			result = append(result, "ok")
		}
	}
	return result, nil
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
