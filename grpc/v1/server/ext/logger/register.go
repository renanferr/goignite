package gigrpclogger

import (
	"context"

	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() {
		return nil
	}

	return []grpc.ServerOption{
		grpc.ChainStreamInterceptor(streamInterceptor()),
		grpc.ChainUnaryInterceptor(unaryInterceptor()),
	}
}
