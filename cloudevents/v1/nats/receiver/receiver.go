package receiver

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
)

func StartReceiver(ctx context.Context, fn interface{}, options *Options) {

	logger := gilog.FromContext(ctx)

	for _, subject := range options.Subjects {

		ct, cl := context.WithCancel(ctx)

		go func(subject string, ctx context.Context, cancel context.CancelFunc) {

			c, err := NewClient(options.Url, subject)
			if err != nil {
				logger.Fatalf("failed to create client: %s", err.Error())
			}

			logger.Infof("connected to the %s with the subject %s", options.Url, subject)

			if err := c.StartReceiver(ctx, fn); err != nil {
				logger.Fatalf("failed to start receiver: %s", err.Error())
			}

			cancel()

		}(subject, ct, cl)

	}

	<-ctx.Done()
}

func StartDefaultReceiver(ctx context.Context, fn interface{}) {
	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	StartReceiver(ctx, fn, o)
}
