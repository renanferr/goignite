package health

import (
	"context"

	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-redis/redis/v7"
)

type ClusterIntegrator struct {
	options *Options
}

func NewClusterIntegrate(options *Options) *ClusterIntegrator {
	return &ClusterIntegrator{options: options}
}

func NewDefaultClusterIntegrator() *ClusterIntegrator {
	o, err := DefaultOptions()
	if err != nil {
		gilog.Fatalf(err.Error())
	}

	return NewClusterIntegrate(o)
}

func (i *ClusterIntegrator) Integrate(ctx context.Context, client *redis.ClusterClient) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis with health")

	checker := NewClusterClientChecker(client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("redis integrated on health with success")

	return nil
}
