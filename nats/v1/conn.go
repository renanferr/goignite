package ginats

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/nats-io/nats.go"
)

type Ext func(context.Context, *nats.Conn) error

func NewConnection(ctx context.Context, options *Options, exts ...Ext) (*nats.Conn, error) {

	logger := gilog.FromContext(ctx)

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

	for _, ext := range exts {
		if err := ext(ctx, conn); err != nil {
			panic(err)
		}
	}

	logger.Infof("Connected to NATS server: %s", options.Url)

	return conn, nil
}

func NewDefaultConnection(ctx context.Context, exts ...Ext) (*nats.Conn, error) {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewConnection(ctx, o, exts...)
}

func disconnectedErrHandler(nc *nats.Conn, err error) {
	gilog.Errorf("Disconnected due to:%s, will attempt reconnects for %.0fm", err)
}

func reconnectedHandler(nc *nats.Conn) {
	gilog.Warnf("Reconnected [%s]", nc.ConnectedUrl())
}

func closedHandler(nc *nats.Conn) {
	gilog.Errorf("Exiting: %v", nc.LastError())
}
