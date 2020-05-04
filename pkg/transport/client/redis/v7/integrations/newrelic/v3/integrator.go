package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/client/redis/v7"
	r "github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

type Integrator struct {
	options *Options
}

func NewIntegrator(options *Options) redis.Integrator {
	return &Integrator{options: options}
}

func (i *Integrator) Integrate(ctx context.Context, client *r.Client) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating r with newrelic")

	if i.options.Enabled {
		client.AddHook(nrredis.NewHook(client.Options()))
	}

	logger.Debug("mongodb integrated with r with success")

	return nil
}
