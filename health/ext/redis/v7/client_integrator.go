package gihealthredis

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	giredis "github.com/b2wdigital/goignite/redis/v7"
	"github.com/go-redis/redis/v7"
)

type ClientIntegrator struct {
	options *Options
}

func ClientIntegrate(options *Options) error {
	integrator := &ClientIntegrator{options: options}
	return gieventbus.Subscribe(giredis.TopicClient, integrator.Integrate)
}

func (i *ClientIntegrator) Integrate(client *redis.Client) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating redis with health")

	checker := NewClientChecker(client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("redis integrated on health with success")

	return nil
}
