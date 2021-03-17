package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/go-redis/redis/v8"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in newrelic")

	client.AddHook(NewHook(client.Options()))

	logger.Debug("redis successfully integrated in newrelic")

	return nil
}
