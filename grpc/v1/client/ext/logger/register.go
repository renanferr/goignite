package gigrpclogger

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.DialOption {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)
	logger.Debug("logger interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainStreamInterceptor(streamInterceptor()),
		grpc.WithChainUnaryInterceptor(unaryInterceptor()),
	}
}
