package http

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/cloudevents/sdk-go"
)

func Start(ctx context.Context, fn interface{}, method string) {

	l := gilog.FromContext(ctx)

	port := GetPort()
	path := GetPath()
	contentType := GetContentType()

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(port),
		cloudevents.WithPath(path),
		cloudevents.WithMethod(method),
	)
	if err != nil {
		l.Fatalf("failed to create transport: %s", err.Error())
	}
	c, err := cloudevents.NewClient(t,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
		cloudevents.WithDataContentType(contentType),
	)
	if err != nil {
		l.Fatalf("failed to create client: %s", err.Error())
	}

	l.Infof("listening on %d:%s", port, path)

	if err := c.StartReceiver(ctx, fn); err != nil {
		l.Fatalf("failed to start receiver: %s", err.Error())
	}

	<-ctx.Done()
}
