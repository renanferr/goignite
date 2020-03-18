package nats

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/nats-io/nats.go"
)

func NewConnection(ctx context.Context, options *Options) (*nats.Conn, error) {

	l := log.FromContext(ctx)

	conn, err := nats.Connect(
		options.Url,
		nats.MaxReconnects(options.MaxReconnects),
		nats.ReconnectWait(options.ReconnectWait),
		nats.DisconnectErrHandler(disconnectedErrHandler),
		nats.ReconnectHandler(reconnectedHandler),
		nats.ClosedHandler(closedHandler),
	)

	if err != nil {
		return nil, err
	}

	if options.Health.Enabled {
		configureHealthCheck(conn, options)
	}

	l.Infof("Connected to NATS server: %s", options.Url)

	return conn, nil
}

func NewDefaultConnection(ctx context.Context) (*nats.Conn, error) {

	l := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewConnection(ctx, o)
}

func configureHealthCheck(conn *nats.Conn, o *Options) {
	cc := NewConnectionChecker(conn)
	hc := health.NewHealthChecker("nats", o.Health.Description, cc, o.Health.Required)

	health.Add(hc)
}

func disconnectedErrHandler(nc *nats.Conn, err error) {
	log.Errorf("Disconnected due to:%s, will attempt reconnects for %.0fm", err)
}

func reconnectedHandler(nc *nats.Conn) {
	log.Warnf("Reconnected [%s]", nc.ConnectedUrl())
}

func closedHandler(nc *nats.Conn) {
	log.Errorf("Exiting: %v", nc.LastError())
}
