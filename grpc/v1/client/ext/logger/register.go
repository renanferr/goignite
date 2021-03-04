package gigrpclogger

import (
	"context"

	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.DialOption {

	if !IsEnabled() {
		return nil
	}

	return []grpc.DialOption{
		grpc.WithChainStreamInterceptor(streamInterceptor()),
		grpc.WithChainUnaryInterceptor(unaryInterceptor()),
	}
}
