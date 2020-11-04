package receiver

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
)

func StartReceiver(ctx context.Context, fn interface{}, options *Options) {

	l := gilog.FromContext(ctx)

	for _, subject := range options.Subjects {

		ct, cl := context.WithCancel(ctx)

		go func(subject string, ctx context.Context, cancel context.CancelFunc) {

			c, err := NewClient(options.Url, subject)
			if err != nil {
				l.Fatalf("failed to create client: %s", err.Error())
			}

			l.Infof("connected to the %s with the subject %s", options.Url, subject)

			if err := c.StartReceiver(ctx, fn); err != nil {
				l.Fatalf("failed to start receiver: %s", err.Error())
			}

			cancel()

		}(subject, ct, cl)

	}

	<-ctx.Done()
}

func StartDefaultReceiver(ctx context.Context, fn interface{}) {
	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	StartReceiver(ctx, fn, o)
}
