package main

import (
	"context"
	"github.com/klpx/talk-golang-context/pkg/auth"
	"github.com/klpx/talk-golang-context/pkg/ctxlog"
	"github.com/klpx/talk-golang-context/pkg/log"
	"github.com/klpx/talk-golang-context/pkg/metrics"
	"github.com/klpx/talk-golang-context/pkg/storage"
	"github.com/klpx/talk-golang-context/pkg/tracing"
	"github.com/klpx/talk-golang-context/pkg/webapp"
	"github.com/redis/go-redis/v9"
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisClient.AddHook(&metrics.RedisMetrics{})

	visitStorage := storage.Make(redisClient)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", webapp.HelloWorld(visitStorage))
	mux.HandleFunc("/status", webapp.ServersStatus())

	handler := tracing.RequestIDMiddleware(mux)
	handler = auth.Middleware(handler)

	logger.Infoc(context.TODO(), "Starting HTTP server on port 3333")
	server := &http.Server{
		Addr:    ":3333",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
