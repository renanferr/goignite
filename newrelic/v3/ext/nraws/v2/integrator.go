package ginraws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	giaws "github.com/b2wdigital/goignite/aws/v2"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	// "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
)

type Integrator struct {
}

func Integrate() error {
	if !IsEnabled() {
		return nil
	}

	integrator := &Integrator{}
	return gieventbus.SubscribeOnce(giaws.TopicConfig, integrator.Integrate)
}

func (i *Integrator) Integrate(cfg *aws.Config) error {

	// logger := gilog.WithTypeOf(*i)
	// logger.Trace("integrating aws with newrelic")
	// nrawssdk.InstrumentHandlers(&cfg.Handlers)
	// logger.Debug("aws integrated with newrelic with success")
	// return nil

	panic("the newrelic nrawssdk-v2 is not compatible with new aws sdk")
}
