package http

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	cloudeventsnats "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/nats"
)

func Start(ctx context.Context, fn interface{}) {

	l := log.FromContext(ctx)

	server := GetServer()
	subject := GetSubject()

	t, err := cloudeventsnats.New(server, subject)
	if err != nil {
		l.Fatalf("failed to create transport: %s", err.Error())
	}

	c, err := client.New(t)
	if err != nil {
		l.Fatalf("failed to create client: %s", err.Error())
	}

	l.Infof("connected to the %s server with the subject %s", server, subject)

	if err := c.StartReceiver(ctx, fn); err != nil {
		l.Fatalf("failed to start receiver: %s", err.Error())
	}

	<-ctx.Done()
}
