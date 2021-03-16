package ginewrelic

import (
	"context"
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var app *newrelic.Application

func Application() *newrelic.Application {
	if app == nil {
		var err error
		if app, err = NewApplication(context.Background()); err != nil {
			panic(err)
		}
	}
	return app
}

func NewApplication(ctx context.Context) (*newrelic.Application, error) {
	logger := gilog.FromContext(ctx)

	enabled := giconfig.Bool(enabled)
	appName := giconfig.String(appName)
	a, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(giconfig.String(license)),
		newrelic.ConfigEnabled(enabled),
		newrelic.ConfigDistributedTracerEnabled(giconfig.Bool(tracerEnabled)),
		newrelic.ConfigLogger(NewLogger()),
		// newrelic.ConfigDebugLogger(log.GetLogger().Output()),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.IgnoreStatusCodes = giconfig.Ints(errorCollectorIgnoreStatusCodes)
			cfg.Labels = giconfig.StringMap(labels)
			cfg.ServerlessMode.Enabled = giconfig.Bool(serverlessModeEnabled)
			cfg.ServerlessMode.AccountID = giconfig.String(serverlessModeAccountID)
			cfg.ServerlessMode.TrustedAccountKey = giconfig.String(serverlessModeTrustedAccountKey)
			cfg.ServerlessMode.PrimaryAppID = giconfig.String(serverlessModePrimaryAppID)
			if apdex, err := time.ParseDuration(giconfig.String(serverlessModeApdexThreshold) + "s"); nil == err {
				cfg.ServerlessMode.ApdexThreshold = apdex
			}
		},
	)

	if err != nil {
		return nil, err
	}

	if enabled {
		logger.Debugf("started a new NewRelic application: %s", appName)
	}

	app = a

	return app, nil
}
