package health

import (
	"context"
	"database/sql"

	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
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

func (i *Integrator) Integrate(ctx context.Context, db *sql.DB) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating godror with health")

	checker := NewChecker(db)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("godror integrated on health with success")

	return nil
}
