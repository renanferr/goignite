package ginrredis

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	ginrredis "github.com/b2wdigital/goignite/newrelic/v3/ext/nrredis"
	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

func Integrate(ctx context.Context, client *redis.Client) error {

	if !ginrredis.IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating redis with newrelic")

	client.AddHook(nrredis.NewHook(client.Options()))
	logger.Debug("redis integrated with newrelic with success")

	return nil
}
