package realip

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Register(ctx context.Context, instance *chi.Mux) error {
	if IsEnabled() {
		instance.Use(middleware.RealIP)
	}

	return nil
}
