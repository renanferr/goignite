package gihealthnats

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	ginats "github.com/b2wdigital/goignite/nats/v1"
	"github.com/nats-io/nats.go"
)

type Integrator struct {
	options *Options
}

func Integrate(options *Options) error {
	integrator := &Integrator{options: options}
	return gieventbus.Subscribe(ginats.TopicConn, integrator.Integrate)
}

func (i *Integrator) Integrate(conn *nats.Conn) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating nats with health")

	checker := NewChecker(conn)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("nats integrated on health with success")

	return nil
}
