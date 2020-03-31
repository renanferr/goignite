package receiver

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/cloudevents/sdk-go/v2/client"
	natsce "github.com/cloudevents/sdk-go/v2/protocol/nats"
	"github.com/nats-io/nats.go"
)

func StartConsumer(ctx context.Context, conn *nats.Conn, fn interface{}, options *Options) {

	l := log.FromContext(ctx)

	for _, subject := range options.Subjects {

		ct, cl := context.WithCancel(ctx)

		go func(subject string, ctx context.Context, cancel context.CancelFunc) {

			c, err := natsce.NewConsumerFromConn(conn, subject)
			if err != nil {
				l.Fatalf("failed to create client: %s", err.Error())
			}

			l.Infof("connected in subject %s", subject)

			cli, err := client.New(c)
			if err != nil {
				l.Fatalf("failed to create client, %s", err.Error())
			}

			if err := cli.StartReceiver(ctx, fn); err != nil {
				l.Fatalf("failed to start receiver: %s", err.Error())
			}

			cancel()

		}(subject, ct, cl)

	}

	<-ctx.Done()
}

func StartDefaultConsumer(ctx context.Context, conn *nats.Conn, fn interface{}) {
	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	StartConsumer(ctx, conn, fn, o)
}
