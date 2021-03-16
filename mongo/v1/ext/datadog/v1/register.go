package gimongodatadog

import (
	"context"

	gidatadog "github.com/b2wdigital/goignite/v2/datadog/v1"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gimongo "github.com/b2wdigital/goignite/v2/mongo/v1"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

func Register(ctx context.Context, conn *gimongo.Conn) error {

	if !IsEnabled() || !gidatadog.IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating mongo in datadog")

	conn.ClientOptions.SetMonitor(mongotrace.NewMonitor())

	logger.Debug("mongo successfully integrated in datadog")

	return nil
}
