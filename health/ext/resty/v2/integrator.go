package gihealthresty

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	giresty "github.com/b2wdigital/goignite/resty/v2"
	"github.com/go-resty/resty/v2"
)

type Integrator struct {
	options *Options
}

func Integrate(options *Options) error {
	integrator := &Integrator{options: options}
	return gieventbus.Subscribe(giresty.TopicClient, integrator.Integrate)
}

func (i *Integrator) Integrate(client *resty.Client) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating resty with health")

	checker := NewChecker(client, i.options)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("resty integrated on health with success")

	return nil
}
