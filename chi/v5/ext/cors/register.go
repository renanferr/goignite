package cors

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Register(ctx context.Context, instance *chi.Mux) error {

	if !isEnabled() {
		return nil
	}

	instance.Use(cors.Handler(cors.Options{
		AllowedOrigins:   getAllowedOrigins(),
		AllowedMethods:   getAllowedMethods(),
		AllowedHeaders:   getAllowedHeaders(),
		AllowCredentials: getAllowedCredentials(),
		ExposedHeaders:   getExposedHeaders(),
		MaxAge:           getMaxAge(),
	}))

	return nil
}
