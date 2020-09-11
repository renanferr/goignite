package ginewrelic

import (
	"context"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var app *newrelic.Application

func Application() *newrelic.Application {
	return app
}

func NewApplication(ctx context.Context) (*newrelic.Application, error) {
	l := gilog.FromContext(ctx)

	enabled := giconfig.Bool(Enabled)
	appName := giconfig.String(AppName)
	a, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(giconfig.String(License)),
		newrelic.ConfigEnabled(enabled),
		newrelic.ConfigDistributedTracerEnabled(giconfig.Bool(TracerEnabled)),
 		newrelic.ConfigLogger(NewLogger()),
 		// newrelic.ConfigDebugLogger(log.GetLogger().Output()),
 		func(cfg *newrelic.Config) {
 			cfg.ErrorCollector.IgnoreStatusCodes = giconfig.Ints(ErrorCollectorIgnoreStatusCodes)
 			cfg.Labels = giconfig.StringMap(Labels)
 			cfg.ServerlessMode.Enabled = giconfig.Bool(ServerlessModeEnabled)
			cfg.ServerlessMode.AccountID = giconfig.String(ServerlessModeAccountID)
			cfg.ServerlessMode.TrustedAccountKey = giconfig.String(ServerlessModeTrustedAccountKey)
			cfg.ServerlessMode.PrimaryAppID = giconfig.String(ServerlessModePrimaryAppID)
			if apdex, err := time.ParseDuration(giconfig.String(ServerlessModeApdexThreshold) + "s"); nil == err {
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
