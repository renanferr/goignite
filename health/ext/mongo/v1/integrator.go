package gihealthmongo

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
	"go.mongodb.org/mongo-driver/mongo"
)

type Integrator struct {
	options *Options
}

func Integrate(options *Options) error {
	integrator := &Integrator{options: options}
	return gieventbus.Subscribe(gimongo.TopicClient, integrator.Integrate)
}

func (i *Integrator) Integrate(client *mongo.Client) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating mongo with health")

	checker := NewChecker(client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("mongo integrated on health with success")

	return nil
}
