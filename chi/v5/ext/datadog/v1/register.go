package datadog

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/chi/v5"
	"github.com/b2wdigital/goignite/v2/datadog/v1"
	"github.com/b2wdigital/goignite/v2/log"
)

func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() || !datadog.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			chiMiddleware(WithServiceName(datadog.Service())),
		},
	}, nil
}
