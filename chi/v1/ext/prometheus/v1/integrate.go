package prometheus

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Integrate(ctx context.Context, instance *chi.Mux) error {

	if !isEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating chi with prometheus")

	prometheusRoute := getRoute()

	logger.Infof("configuring prometheus metrics router on %s", prometheusRoute)
	instance.Handle(prometheusRoute, promhttp.Handler())

	instance.Use(promMiddleware)

	logger.Debug("prometheus integrated with echo with success")

	return nil
}
