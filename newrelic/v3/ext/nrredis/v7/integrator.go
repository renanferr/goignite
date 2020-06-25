package ginrredis

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	giredis "github.com/b2wdigital/goignite/redis/v7"
	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

type Integrator struct {
}

func Integrate() error {
	integrator := &Integrator{}
	return gieventbus.SubscribeOnce(giredis.TopicClient, integrator.Integrate)
}

func (i *Integrator) Integrate(client *redis.Client) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating redis with newrelic")

	if IsEnabled() {
		client.AddHook(nrredis.NewHook(client.Options()))
		logger.Debug("redis integrated with newrelic with success")
	} else {
		logger.Debug("redis integration is disabled")
	}

	return nil
}
