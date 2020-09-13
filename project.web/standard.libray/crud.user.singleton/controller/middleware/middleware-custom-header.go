// Go Api server
// @jeffotoni
// 2019-01-08

package middleware

import (
	"net/http"
	"strings"
)

// CustomHeaders adds cache and basic HTTP headers to response
func CustomHeaders() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Server", "api.crud.user")
			w.Header().Add("Content-Type", "application/json")
			//w.Header().Add(runtime.HeaderContentType, "application/json")

			if strings.Contains(r.URL.RawQuery, "seed") {
				w.Header().Add("Expires", "36000")
			} else {
				//w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
				w.Header().Add("Cache-Control", "public")
				w.Header().Add("Cache-Control", "max-age=31536000")
				//w.Header().Add("Expires", "0")
				//w.Header().Add("Pragma", "no-cache")
				//w.Header().Add("Expires", "0")
			}

			h.ServeHTTP(w, r)
		})
	}
}
