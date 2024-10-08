package webapp

import (
	"context"
	"errors"
	"fmt"
	"github.com/klpx/talk-golang-context/pkg/auth"
	"github.com/klpx/talk-golang-context/pkg/ctxlog"
	"github.com/klpx/talk-golang-context/pkg/srvstatus"
	"github.com/klpx/talk-golang-context/pkg/storage"
	"io"
	"net/http"
	"time"
)

func HelloWorld(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var userCount, totalCount int64
		query := r.URL.Query()
		name := query.Get("name")
		if name != "" {
			ctxlog.Infof(ctx, "user %s visited us", name)
			userCount, totalCount, _ = store.RecordVisit(ctx, name)
			_, _ = io.WriteString(w, fmt.Sprintf("Hello, %s! This is your %d visit.\n", name, userCount))
		} else {
			name = "World"
			userCount, totalCount, _ = store.RecordVisit(ctx, "anonymous")
			_, _ = io.WriteString(w, fmt.Sprintf("Hello, world!\n"))
		}
		_, _ = io.WriteString(w, fmt.Sprintf("This site was visited %d times!\n", totalCount))
	}
}

func ServersStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer func() {
			cancel()
		}()

		statuses, err := srvstatus.CheckServers(ctx)
		if err != nil {
			if errors.Is(err, auth.ErrNotAuthorized) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("forbidden"))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}

		for i, status := range statuses {
			_, _ = io.WriteString(w, fmt.Sprintf("server[%d] = %s\n", i, status))
		}
	}
}
