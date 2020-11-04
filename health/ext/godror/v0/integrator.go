package gihealthgodror

import (
	"database/sql"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gigodror "github.com/b2wdigital/goignite/godror/v0"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
)

type Integrator struct {
	options *Options
}

func Integrate(options *Options) error {
	integrator := &Integrator{options: options}
	return gieventbus.Subscribe(gigodror.TopicDB, integrator.Integrate)
}

func (i *Integrator) Integrate(db *sql.DB) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating godror with health")

	checker := NewChecker(db)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("godror integrated on health with success")

	return nil
}
