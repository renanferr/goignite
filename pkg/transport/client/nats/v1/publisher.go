package nats

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Publisher struct {
	conn    *nats.Conn
	options *Options
}

func NewPublisher(ctx context.Context, options *Options) (*Publisher, error) {
	conn, err := NewConnection(ctx, options)
	if err != nil {
		return nil, err
	}
	return &Publisher{conn, options}, nil
}

func NewDefaultPublisher(ctx context.Context) (*Publisher, error) {
	options, err := DefaultOptions()
	if err != nil {
		return nil, err
	}
	return NewPublisher(ctx, options)
}

func (p *Publisher) Publish(ctx context.Context, msg *nats.Msg) error {

	if p.options.NewRelic.Enabled {
		txn := newrelic.FromContext(ctx)
		seg := nrnats.StartPublishSegment(txn, p.conn, msg.Subject)
		defer seg.End()
	}

	return p.conn.PublishMsg(msg)
}

func (p *Publisher) Conn() *nats.Conn {
	return p.conn
}
