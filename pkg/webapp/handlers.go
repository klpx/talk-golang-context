package webapp

import (
	"fmt"
	"github.com/klpx/talk-golang-context/pkg/ctxlog"
	"io"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()
	name := query.Get("name")
	if name != "" {
		ctxlog.Infof(ctx, "user %s visited us", name)
	} else {
		name = "World"
	}
	_, _ = io.WriteString(w, fmt.Sprintf("Hello, %s!\n", name))
}
