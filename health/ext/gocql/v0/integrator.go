package gihealtgocql

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gigocql "github.com/b2wdigital/goignite/gocql/v0"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/gocql/gocql"
)

type Integrator struct {
	options *Options
}

func Integrate(options *Options) error {
	integrator := &Integrator{options: options}
	return gieventbus.Subscribe(gigocql.TopicSession, integrator.Integrate)
}

func (i *Integrator) Integrate(session *gocql.Session) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating gocql with health")

	checker := NewChecker(session)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("gocql integrated on health with success")

	return nil
}
