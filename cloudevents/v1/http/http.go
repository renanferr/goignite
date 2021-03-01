package http

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	cloudevents "github.com/cloudevents/sdk-go"
)

func Start(ctx context.Context, fn interface{}, method string) {

	logger := gilog.FromContext(ctx)

	port := GetPort()
	path := GetPath()
	contentType := GetContentType()

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(port),
		cloudevents.WithPath(path),
		cloudevents.WithMethod(method),
	)
	if err != nil {
		logger.Fatalf("failed to create transport: %s", err.Error())
	}
	c, err := cloudevents.NewClient(t,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
		cloudevents.WithDataContentType(contentType),
	)
	if err != nil {
		logger.Fatalf("failed to create client: %s", err.Error())
	}

	logger.Infof("listening on %d:%s", port, path)

	if err := c.StartReceiver(ctx, fn); err != nil {
		logger.Fatalf("failed to start receiver: %s", err.Error())
	}

	<-ctx.Done()
}
