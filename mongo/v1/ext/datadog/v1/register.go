package datadog

import (
	"context"

	"github.com/b2wdigital/goignite/v2/datadog/v1"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/mongo/v1"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

func Register(ctx context.Context, conn *mongo.Conn) error {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating mongo in datadog")

	conn.ClientOptions.SetMonitor(mongotrace.NewMonitor())

	logger.Debug("mongo successfully integrated in datadog")

	return nil
}
