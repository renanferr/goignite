package newrelic

import (
	"context"

	ginats "github.com/b2wdigital/goignite/nats/v1"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type PublisherMiddleware struct {
}

func (p *PublisherMiddleware) Before(ctx context.Context, conn *nats.Conn, msg *nats.Msg) (context.Context, error) {
	if !isEnabled() {
		return ctx, nil
	}

	txn := newrelic.FromContext(ctx)
	seg := nrnats.StartPublishSegment(txn, conn, msg.Subject)

	return context.WithValue(ctx, "seg", seg), nil
}

func (p *PublisherMiddleware) After(ctx context.Context) error {
	seg := ctx.Value("seg").(*newrelic.MessageProducerSegment)
	seg.End()
	return nil
}

func NewPublisherMiddleware() ginats.PublisherMiddleware {
	return &PublisherMiddleware{}
}
