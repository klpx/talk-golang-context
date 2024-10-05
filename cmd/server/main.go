package main

import (
	"context"
	"github.com/klpx/talk-golang-context/pkg/ctxlog"
	"github.com/klpx/talk-golang-context/pkg/log"
	"github.com/klpx/talk-golang-context/pkg/tracing"
	"github.com/klpx/talk-golang-context/pkg/webapp"
	"net/http"
)

func main() {
	logger := log.MakeLogger(func(ctx context.Context) map[string]string {
		result := map[string]string{}
		if reqID, ok := tracing.CtxRequestID.Value(ctx); ok {
			result["request_id"] = reqID
		}
		return result
	})
	ctxlog.SetGlobalLogger(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", webapp.HelloWorld)

	handler := tracing.RequestIDMiddleware(mux)

	logger.Infoc(context.TODO(), "Starting HTTP server on port 3333")
	err := http.ListenAndServe(":3333", handler)
	if err != nil {
		panic(err)
	}
}
