package nats

import (
	"context"
	"errors"

	"github.com/nats-io/nats.go"
)

type ConnectionChecker struct {
	conn *nats.Conn
}

func (c *ConnectionChecker) Check(ctx context.Context) error {

	var err error

	if !c.conn.IsConnected() {
		err = errors.New("Not connected")
	}

	return err
}

func NewConnectionChecker(conn *nats.Conn) *ConnectionChecker {
	return &ConnectionChecker{conn: conn}
}
