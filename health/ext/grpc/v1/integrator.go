package gihealthgrpc

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gigrpc "github.com/b2wdigital/goignite/grpc/v1/client"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	"google.golang.org/grpc"
)

type Integrator struct {
	options *Options
}

func Integrate(options *Options) error {
	integrator := &Integrator{options: options}
	return gieventbus.Subscribe(gigrpc.TopicClientConn, integrator.Integrate)
}

func (i *Integrator) Integrate(conn *grpc.ClientConn) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating grpc with health")

	checker := NewChecker(conn)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("grpc integrated on health with success")

	return nil
}
