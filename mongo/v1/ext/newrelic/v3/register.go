package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/mongo/v1"
	"github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
)

func Register(ctx context.Context, conn *mongo.Conn) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating mongo in newrelic")

	nrMon := nrmongo.NewCommandMonitor(nil)

	conn.ClientOptions.SetMonitor(nrMon)

	logger.Debug("mongo successfully integrated in newrelic")

	return nil
}
