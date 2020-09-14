package giresty

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	giresty "github.com/b2wdigital/goignite/resty/v2"
	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Integrator struct {
}

func Integrate() error {
	integrator := &Integrator{}
	return gieventbus.Subscribe(giresty.TopicClient, integrator.Integrate)
}

func (i *Integrator) Integrate(client *resty.Client) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating resty with newrelic")

	if IsEnabled() {

		client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {

			logger := gilog.FromContext(request.Context())

			txn := newrelic.FromContext(request.Context())

			txn.InsertDistributedTraceHeaders(request.Header)

			logger.Debugf("rest request processing")

			return nil
		})

		logger.Debug("resty integrated with newrelic with success")
	} else {
		logger.Debug("resty integration is disabled")
	}

	return nil
}
