package recoverer

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Register(ctx context.Context, instance *chi.Mux) error {
	if IsEnabled() {
		instance.Use(middleware.Recoverer)
	}

	return nil
}
