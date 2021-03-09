package ginatsnewrelic

import (
	"context"

	ginats "github.com/b2wdigital/goignite/v2/nats/v1"
	ginewrelic "github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type PublisherRegister struct {
}

func (p *PublisherRegister) Before(ctx context.Context, conn *nats.Conn, msg *nats.Msg) (context.Context, error) {
	if !IsEnabled() {
		return ctx, nil
	}

	txn := ginewrelic.FromContext(ctx)
	seg := nrnats.StartPublishSegment(txn, conn, msg.Subject)

	return context.WithValue(ctx, "seg", seg), nil
}

func (p *PublisherRegister) After(ctx context.Context) error {
	seg := ctx.Value("seg").(*newrelic.MessageProducerSegment)
	seg.End()
	return nil
}

func NewPublisherRegister() ginats.PublisherMiddleware {
	return &PublisherRegister{}
}
