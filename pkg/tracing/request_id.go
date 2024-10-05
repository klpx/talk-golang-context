package tracing

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/klpx/talk-golang-context/pkg/ctxstore"
	"net/http"
)

var CtxRequestID = ctxstore.MakeStore[string]()

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = randID()
		}
		w.Header().Add("X-Request-ID", reqID)
		newCtx := CtxRequestID.WithValue(r.Context(), reqID)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func randID() string {
	rBytes := make([]byte, 12)
	_, _ = rand.Read(rBytes)
	return base64.URLEncoding.EncodeToString(rBytes)
}
