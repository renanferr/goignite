package gihealthelasticsearch

import (
	gielasticsearch "github.com/b2wdigital/goignite/elasticsearch/v8"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/elastic/go-elasticsearch/v8"
)

type Integrator struct {
	options *Options
}

func Integrate(options *Options) error {
	integrator := &Integrator{options: options}
	return gieventbus.Subscribe(gielasticsearch.TopicClient, integrator.Integrate)
}

func (i *Integrator) Integrate(client *elasticsearch.Client) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating elasticsearch with health")

	checker := NewChecker(client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("elasticsearch integrated on health with success")

	return nil
}
