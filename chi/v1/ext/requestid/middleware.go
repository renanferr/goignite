package requestid

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Middleware(ctx context.Context, instance *chi.Mux) error {
	if isEnabled() {
		instance.Use(middleware.RequestID)
	}

	return nil
}
