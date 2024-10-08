package auth

import (
	"errors"
	"github.com/klpx/talk-golang-context/pkg/ctxstore"
	"net/http"
	"strings"
)

////

type Principal struct {
	authorized bool
}

func (p *Principal) IsAuthorized() bool {
	if p == nil {
		return false
	}
	return p.authorized
}

////

var Context = ctxstore.MakeStore[*Principal]()

////

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if token, ok := strings.CutPrefix(authHeader, "Bearer "); ok && token == "token" {
			newCtx := Context.WithValue(r.Context(), &Principal{authorized: true})
			next.ServeHTTP(w, r.WithContext(newCtx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

var ErrNotAuthorized = errors.New("not authorized")
