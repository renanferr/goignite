package receiver

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	natsce "github.com/cloudevents/sdk-go/protocol/nats/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/nats-io/nats.go"
)

func StartConsumer(ctx context.Context, conn *nats.Conn, fn interface{}, options *Options) {

	logger := gilog.FromContext(ctx)

	for _, subject := range options.Subjects {

		ct, cl := context.WithCancel(ctx)

		go func(subject string, ctx context.Context, cancel context.CancelFunc) {

			c, err := natsce.NewConsumerFromConn(conn, subject)
			if err != nil {
				logger.Fatalf("failed to create client: %s", err.Error())
			}

			logger.Infof("connected in subject %s", subject)

			cli, err := client.New(c)
			if err != nil {
				logger.Fatalf("failed to create client, %s", err.Error())
			}

			if err := cli.StartReceiver(ctx, fn); err != nil {
				logger.Fatalf("failed to start receiver: %s", err.Error())
			}

			cancel()

		}(subject, ct, cl)

	}

	<-ctx.Done()
}

func StartDefaultConsumer(ctx context.Context, conn *nats.Conn, fn interface{}) {
	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	StartConsumer(ctx, conn, fn, o)
}
