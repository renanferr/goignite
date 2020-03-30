package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

func NewClient(ctx context.Context) (*nr.Application, error) {
	l := log.FromContext(ctx)

	enabled := config.Bool(NewRelicEnabled)
	appName := config.String(NewRelicAppName)
	app, err := nr.NewApplication(
		nr.ConfigAppName(appName),
		nr.ConfigLicense(config.String(NewRelicLicense)),
		nr.ConfigEnabled(enabled),
		nr.ConfigDistributedTracerEnabled(config.Bool(NewRelicDistributedTracerEnabled)),
	)

	if err != nil {
		return nil, err
	}

	if enabled {
		l.Infof("started a new NewRelic application: %s", appName)
	}

	return app, nil
}
