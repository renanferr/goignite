package cors

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Register(ctx context.Context, instance *chi.Mux) error {
	if !IsEnabled() {
		return nil
	}

	instance.Use(cors.Handler(cors.Options{
		AllowedOrigins:   GetAllowedOrigins(),
		AllowedMethods:   GetAllowedMethods(),
		AllowedHeaders:   GetAllowedHeaders(),
		AllowCredentials: GetAllowedCredentials(),
		ExposedHeaders:   GetExposedHeaders(),
		MaxAge:           GetMaxAge(),
	}))

	return nil
}
