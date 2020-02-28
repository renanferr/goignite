package nats

import (
	"context"

	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/health"
	"github.com/jpfaria/goignite/pkg/log/logrus"
	h "github.com/jpfaria/goignite/pkg/pubsub/client/nats/health"
	"github.com/jpfaria/goignite/pkg/pubsub/client/nats/model"
	"github.com/nats-io/nats.go"
)

func NewClient(ctx context.Context, options model.Options) (*nats.Conn, error) {

	log := logrus.FromContext(ctx)

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

	log.Infof("Connected to NATS server: %s", options.Url)

	return conn, nil
}

func NewDefaultClient(ctx context.Context) (*nats.Conn, error) {

	log := logrus.FromContext(ctx)

	o := model.Options{}

	err := config.UnmarshalWithPath("pubsub.client.nats", &o)
	if err != nil {
		log.Fatal(err)
	}

	return NewClient(ctx, o)

}

func configureHealthCheck(conn *nats.Conn, o model.Options) {
	cc := h.NewClientChecker(conn)
	hc := health.NewHealthChecker("nats", o.Health.Description, cc, o.Health.Required)

	health.Add(hc)
}
