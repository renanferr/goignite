package http

import (
	"context"

	"github.com/cloudevents/sdk-go"
	"github.com/jpfaria/goignite/pkg/log/logrus"
	"github.com/jpfaria/goignite/pkg/serverless/cloudevents/transport/http/config"
)

func Start(ctx context.Context, fn interface{}) {

	log := logrus.FromContext(ctx)

	port := config.GetPort()
	path := config.GetPath()

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(port),
		cloudevents.WithPath(path),
	)
	if err != nil {
		log.Fatalf("failed to create transport: %s", err.Error())
	}
	c, err := cloudevents.NewClient(t,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err.Error())
	}

	log.Infof("listening on :%d%s", port, path)

	if err := c.StartReceiver(ctx, fn); err != nil {
		log.Fatalf("failed to start receiver: %s", err.Error())
	}

	<-ctx.Done()
}
