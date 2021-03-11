package gichitid

import (
	"context"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	"github.com/b2wdigital/goignite/v2/info"
	gilog "github.com/b2wdigital/goignite/v2/log"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context) (*gichi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("enabling tid middleware in chi")

	return &gichi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			tidMiddleware(),
		},
	}, nil
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
