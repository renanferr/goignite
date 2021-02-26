package gihealthredis

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	giredis "github.com/b2wdigital/goignite/redis/v8"
	"github.com/go-redis/redis/v8"
)

type ClusterIntegrator struct {
	options *Options
}

func ClusterIntegrate(options *Options) error {
	integrator := &ClusterIntegrator{options: options}
	return gieventbus.Subscribe(giredis.TopicClusterClient, integrator.Integrate)
}

func (i *ClusterIntegrator) Integrate(client *redis.ClusterClient) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating redis with health")

	checker := NewClusterClientChecker(client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("redis integrated on health with success")

	return nil
}
