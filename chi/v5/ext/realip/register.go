package gichirealip

import (
	"context"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-chi/chi/v5/middleware"
)

func Register(ctx context.Context) (*gichi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("enabling realIP middleware in chi")

	return &gichi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			middleware.RealIP,
		},
	}, nil
}
