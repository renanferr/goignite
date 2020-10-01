package ginrmongo

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Integrator struct {
}

func Integrate() error {
	if !IsEnabled() {
		return nil
	}

	integrator := &Integrator{}
	return gieventbus.SubscribeOnce(gimongo.TopicClientOptions, integrator.Integrate)
}

func (i *Integrator) Integrate(clientOptions *options.ClientOptions) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating mongodb with newrelic")

	nrMon := nrmongo.NewCommandMonitor(nil)

	clientOptions.SetMonitor(nrMon)
	logger.Debug("mongodb integrated with newrelic with success")

	return nil
}
