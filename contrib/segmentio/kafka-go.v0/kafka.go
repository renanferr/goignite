package kafka

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/segmentio/kafka-go"
)

func NewLeaderConnection(ctx context.Context, o *Options) (conn *kafka.Conn, err error) {

	logger := log.FromContext(ctx)

	conn, err = kafka.DialLeader(context.Background(), o.Network, o.Address, o.Topic, o.Partition)
	if err != nil {
		logger.Fatal("failed to dial leader:", err)
	}

	logger.Infof("Created kafka connection to %v", o.Address)

	return conn, err

}

func NewDefaultLeaderConnection(ctx context.Context) (*kafka.Conn, error) {

	logger := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewLeaderConnection(ctx, o)
}
