package gichiprometheus

import (
	"context"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(ctx context.Context) (*gichi.Config, error) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("configuring prometheus")

	prometheusRoute := getRoute()

	return &gichi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			promMiddleware,
		},
		Handlers: []gichi.ConfigHandler{
			{
				Handler: promhttp.Handler(),
				Pattern: prometheusRoute,
			},
		},
	}, nil

}
