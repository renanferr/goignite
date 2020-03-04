package http

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/b2wdigital/goignite/pkg/serverless/cloudevents/transport/http/config"
	"github.com/cloudevents/sdk-go"
)

func Start(ctx context.Context, fn interface{}, method string) {

	log := logrus.FromContext(ctx)

	port := config.GetPort()
	path := config.GetPath()
	contentType := config.GetContentType()

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(port),
		cloudevents.WithPath(path),
		cloudevents.WithMethod(method),
	)
	if err != nil {
		log.Fatalf("failed to create transport: %s", err.Error())
	}
	c, err := cloudevents.NewClient(t,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
		cloudevents.WithDataContentType(contentType),
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
