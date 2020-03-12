package nats

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/nats-io/nats.go"
)

func NewClient(ctx context.Context, options Options) (*nats.Conn, error) {

	l := log.FromContext(ctx)

	conn, err := nats.Connect(
		options.Url,
		nats.MaxReconnects(options.MaxReconnects),
		nats.ReconnectWait(options.ReconnectWait))

	if err != nil {
		return nil, err
	}

	// defer conn.Close()

	if options.Health.Enabled {
		configureHealthCheck(conn, options)
	}

	l.Infof("Connected to NATS server: %s", options.Url)

	return conn, nil
}

func NewDefaultClient(ctx context.Context) (*nats.Conn, error) {

	l := log.FromContext(ctx)

	o := Options{}

	err := config.UnmarshalWithPath("transport.client.nats", &o)
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewClient(ctx, o)

}

func configureHealthCheck(conn *nats.Conn, o Options) {
	cc := NewClientChecker(conn)
	hc := health.NewHealthChecker("nats", o.Health.Description, cc, o.Health.Required)

	health.Add(hc)
}
