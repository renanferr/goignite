package ginrmongo

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Integrate(ctx context.Context, clientOptions *options.ClientOptions) error {

	if !isEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating mongodb with newrelic")

	nrMon := nrmongo.NewCommandMonitor(nil)

	clientOptions.SetMonitor(nrMon)
	logger.Debug("mongodb integrated with newrelic with success")

	return nil
}
