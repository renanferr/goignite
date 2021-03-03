package newrelic

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating redis with newrelic")

	client.AddHook(nrredis.NewHook(client.Options()))
	logger.Debug("redis integrated with newrelic with success")

	return nil
}
