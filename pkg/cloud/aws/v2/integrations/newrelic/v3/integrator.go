package newrelic

import (
	"context"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"

	"github.com/b2wdigital/goignite/pkg/cloud/aws/v2"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
)

type Integrator struct {
	options *Options
}

func NewIntegrator(options *Options) aws.Integrator {
	return &Integrator{options: options}
}

func (i *Integrator) Integrate(ctx context.Context, cfg *awssdk.Config) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating mongodb with newrelic")

	if i.options.Enabled {
		nrawssdk.InstrumentHandlers(&cfg.Handlers)
	}

	logger.Debug("mongodb integrated with newrelic with success")

	return nil
}
