package ginraws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
)

func Integrate(ctx context.Context, cfg *aws.Config) error {

	if !IsEnabled() {
		return nil
	}

	// logger := gilog.WithTypeOf(*i)
	// logger.Trace("integrating aws with newrelic")
	// nrawssdk.InstrumentHandlers(&cfg.Handlers)
	// logger.Debug("aws integrated with newrelic with success")
	// return nil

	panic("the newrelic nrawssdk-v2 is not compatible with new aws sdk")
}
