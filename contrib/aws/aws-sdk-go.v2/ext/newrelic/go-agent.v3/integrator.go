package newrelic

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	newrelic "github.com/b2wdigital/goignite/v2/contrib/newrelic/go-agent.v3"
	// "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
)

func Integrate(ctx context.Context, cfg *aws.Config) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	// logger := log.WithTypeOf(*i)
	// logger.Trace("integrating aws with newrelic")
	// nrawssdk.InstrumentHandlers(&cfg.Handlers)
	// logger.Debug("aws integrated with newrelic with success")
	// return nil

	panic("the newrelic nrawssdk-v2 is not compatible with new aws sdk")
}
