package health

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/go-redis/redis/v8"
)

type ClientIntegrator struct {
	options *Options
}

func NewClientIntegrator(options *Options) *ClientIntegrator {
	return &ClientIntegrator{options: options}
}

func NewDefaultClientIntegrator() *ClientIntegrator {
	o, err := DefaultOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewClientIntegrator(o)
}

func (i *ClientIntegrator) Register(ctx context.Context, client *redis.Client) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis in health")

	checker := NewClientChecker(client)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}
