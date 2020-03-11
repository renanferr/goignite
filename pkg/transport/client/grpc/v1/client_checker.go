package v1

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type ClientChecker struct {
	conn *grpc.ClientConn
}

func (c *ClientChecker) Check(ctx context.Context) error {

	var err error

	if c.conn.GetState() != connectivity.Ready {
		err = errors.New("not ready")
	}

	return err
}

func NewClientChecker(conn *grpc.ClientConn) *ClientChecker {
	return &ClientChecker{conn: conn}
}
