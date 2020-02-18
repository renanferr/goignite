package nats

import (
	"context"

	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/logging/logrus"
	"github.com/jpfaria/goignite/pkg/pubsub/nats/model"
	"github.com/nats-io/nats.go"
)

func NewClient(ctx context.Context, options model.Options) (*nats.Conn, error) {

	log := logrus.FromContext(ctx)

	natsConnection, err := nats.Connect(
		options.Url,
		nats.MaxReconnects(options.MaxReconnects),
		nats.ReconnectWait(options.ReconnectWait))

	if err != nil {
		return nil, err
	}

	// defer natsConnection.Close()

	log.Infof("Connected to NATS server: %s", options.Url)

	return natsConnection, nil

}


func NewDefaultClient(ctx context.Context) (*nats.Conn, error) {

	log := logrus.FromContext(ctx)

	o := model.Options{}

	err := config.UnmarshalWithPath("pubsub.nats", &o)
	if err != nil {
		log.Fatal(err)
	}

	return NewClient(ctx, o)

}
