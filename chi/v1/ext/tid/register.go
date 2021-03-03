package tid

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/info"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, instance *chi.Mux) error {
	if IsEnabled() {
		instance.Use(tidMiddleware())
	}

	return nil
}

// tidMiddleware is a middleware that looks for a XTID value inside the http.Request
// and generate one if it does not exists.
func tidMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tid := r.Header.Get("X-TID")
			if tid == "" {
				tid = info.AppName + "-" + uuid.NewV4().String()
			}
			w.Header().Set("X-TID", tid)

			ctx := context.WithValue(r.Context(), "x-tid", tid)
			r.WithContext(ctx)
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
