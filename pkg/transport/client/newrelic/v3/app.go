package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var app *newrelic.Application

func Application() *newrelic.Application {
	return app
}

func NewApplication(ctx context.Context) (*newrelic.Application, error) {
	l := log.FromContext(ctx)

	enabled := config.Bool(Enabled)
	appName := config.String(AppName)
	a, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(config.String(License)),
		newrelic.ConfigEnabled(enabled),
		newrelic.ConfigDistributedTracerEnabled(config.Bool(TracerEnabled)),
		func(cfg *newrelic.Config) { // configure labels
			cfg.Labels = config.StringMap(License)
		},
	)

	if err != nil {
		return nil, err
	}

	if enabled {
		l.Infof("started a new NewRelic application: %s", appName)
	}

	app = a

	return app, nil
}
