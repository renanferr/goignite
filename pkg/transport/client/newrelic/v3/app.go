package newrelic

import (
	"context"
	"time"

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
 		newrelic.ConfigLogger(NewLogger()),
 		// newrelic.ConfigDebugLogger(log.GetLogger().Output()),
 		func(cfg *newrelic.Config) {
 			cfg.Labels = config.StringMap(Labels)
 			cfg.ServerlessMode.Enabled = config.Bool(ServerlessModeEnabled)
			cfg.ServerlessMode.AccountID = config.String(ServerlessModeAccountID)
			cfg.ServerlessMode.TrustedAccountKey = config.String(ServerlessModeTrustedAccountKey)
			cfg.ServerlessMode.PrimaryAppID = config.String(ServerlessModePrimaryAppID)
			if apdex, err := time.ParseDuration(config.String(ServerlessModeApdexThreshold) + "s"); nil == err {
				cfg.ServerlessMode.ApdexThreshold = apdex
			}
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
