package health

import (
	"context"

	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
)

type Integrator struct {
	options *Options
}

func NewIntegrator(options *Options) *Integrator {
	return &Integrator{options: options}
}

func NewDefaultIntegrator() *Integrator {
	o, err := DefaultOptions()
	if err != nil {
		gilog.Fatalf(err.Error())
	}

	return NewIntegrator(o)
}

func (i *Integrator) Integrate(ctx context.Context, conn *gimongo.Conn) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating mongo with health")

	checker := NewChecker(conn.Client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("mongo integrated on health with success")

	return nil
}
