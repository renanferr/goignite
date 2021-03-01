package health

import (
	"context"

	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/gocql/gocql"
)

type Integrator struct {
	options *Options
}

func NewIntegrator(options *Options) *Integrator {
	return &Integrator{options: options}
}

func (i *Integrator) Integrate(ctx context.Context, session *gocql.Session) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating gocql with health")

	checker := NewChecker(session)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("gocql integrated on health with success")

	return nil
}
