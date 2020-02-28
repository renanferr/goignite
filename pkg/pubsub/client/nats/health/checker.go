package health

import (
	"context"
	"errors"

	"github.com/nats-io/nats.go"
)

type ClientChecker struct {
	conn *nats.Conn
}

func (c *ClientChecker) Check(ctx context.Context) error {

	var err error

	if !c.conn.IsConnected() {
		err = errors.New("Not connected")
	}

	return err
}

func NewClientChecker(conn *nats.Conn) *ClientChecker {
	return &ClientChecker{conn: conn}
}
