package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type ClientConnChecker struct {
	conn *grpc.ClientConn
}

func (c *ClientConnChecker) Check(ctx context.Context) error {

	var err error

	if c.conn.GetState() != connectivity.Ready {
		err = errors.New("not ready")
	}

	return err
}

func NewClientConnChecker(conn *grpc.ClientConn) *ClientConnChecker {
	return &ClientConnChecker{conn: conn}
}
