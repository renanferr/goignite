package gielasticsearchhealth

import (
	"context"

	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/elastic/go-elasticsearch/v8"
)

type Integrator struct {
	options *Options
}

func NewIntegrator(options *Options) *Integrator {
	return &Integrator{options: options}
}

func (i *Integrator) Register(ctx context.Context, client *elasticsearch.Client) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating elasticsearch in health")

	checker := NewChecker(client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("elasticsearch successfully integrated in health")

	return nil
}
