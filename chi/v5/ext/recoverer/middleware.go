package recoverer

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Middleware(ctx context.Context, instance *chi.Mux) error {
	if isEnabled() {
		instance.Use(middleware.Recoverer)
	}

	return nil
}
