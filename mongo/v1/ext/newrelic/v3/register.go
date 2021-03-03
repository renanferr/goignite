package newrelic

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
)

func Register(ctx context.Context, conn *gimongo.Conn) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating mongodb with newrelic")

	nrMon := nrmongo.NewCommandMonitor(nil)

	conn.ClientOptions.SetMonitor(nrMon)
	logger.Debug("mongodb integrated with newrelic with success")

	return nil
}
