package gichicors

import (
	"context"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-chi/cors"
)

func Register(ctx context.Context) (*gichi.Config, error) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)
	logger.Tracef("configuring cors")

	return &gichi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			cors.Handler(cors.Options{
				AllowedOrigins:   getAllowedOrigins(),
				AllowedMethods:   getAllowedMethods(),
				AllowedHeaders:   getAllowedHeaders(),
				AllowCredentials: getAllowedCredentials(),
				ExposedHeaders:   getExposedHeaders(),
				MaxAge:           getMaxAge(),
			}),
		},
	}, nil
}
