package logger

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.DialOption {

	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("logger interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainStreamInterceptor(streamInterceptor()),
		grpc.WithChainUnaryInterceptor(unaryInterceptor()),
	}
}
