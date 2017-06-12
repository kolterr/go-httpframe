package algoutil

import (
	"net/http"
	"strings"

	"wukong/internal/cache"
	"wukong/internal/encoding"
	"wukong/internal/errutil"

	"golang.org/x/net/context"
)

func AuthorizationFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.TrimSpace(r.Header.Get("Authorization"))
		if auth == "" {
			encoding.EncodeError(context.Background(), errutil.ErrPermissionDenied, w)
			return
		}

		if !cache.Exists(auth) {
			encoding.EncodeError(context.Background(), errutil.ErrTokenExpired, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
