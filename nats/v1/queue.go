package ginats

import (
	"context"

	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
)

type Queue struct {
	conn    *nats.Conn
	options *Options
}

func NewQueue(ctx context.Context, options *Options) (*Queue, error) {
	conn, err := NewConnection(ctx, options)
	if err != nil {
		return nil, err
	}
	return &Queue{conn, options}, nil
}

func NewDefaultQueue(ctx context.Context) (*Queue, error) {
	options, err := DefaultOptions()
	if err != nil {
		return nil, err
	}
	return NewQueue(ctx, options)
}

func (p *Queue) Subscribe(subj string, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {

	if p.options.NewRelic.Enabled {
		return p.conn.QueueSubscribe(subj, queue, nrnats.SubWrapper(ginewrelic.Application(), cb))
	}

	return p.conn.QueueSubscribe(subj, queue, cb)
}

func (p *Queue) Conn() *nats.Conn {
	return p.conn
}
