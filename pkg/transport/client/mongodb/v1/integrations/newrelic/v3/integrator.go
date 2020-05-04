package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/client/mongodb/v1"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Integrator struct {
	options *Options
}

func NewIntegrator(options *Options) mongodb.Integrator {
	return &Integrator{options: options}
}

func (i *Integrator) Integrate(ctx context.Context, clientOptions *options.ClientOptions) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating mongodb with newrelic")

	nrMon := nrmongo.NewCommandMonitor(nil)

	if i.options.Enabled {
		clientOptions.SetMonitor(nrMon)
	}

	logger.Debug("mongodb integrated with newrelic with success")

	return nil
}
