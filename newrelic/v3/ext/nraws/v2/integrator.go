package ginraws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	giaws "github.com/b2wdigital/goignite/aws/v2"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
)

type Integrator struct {
}

func Integrate() error {
	integrator := &Integrator{}
	return gieventbus.SubscribeOnce(giaws.TopicConfig, integrator.Integrate)
}

func (i *Integrator) Integrate(cfg *aws.Config) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating mongodb with newrelic")

	if IsEnabled() {
		nrawssdk.InstrumentHandlers(&cfg.Handlers)
	}

	logger.Debug("mongodb integrated with newrelic with success")

	return nil
}
