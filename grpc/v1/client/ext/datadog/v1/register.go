package datadog

import (
	"context"

	"github.com/b2wdigital/goignite/v2/datadog/v1"
	"github.com/b2wdigital/goignite/v2/log"
	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.DialOption {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("datadog interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(grpctrace.UnaryClientInterceptor()),
		grpc.WithChainStreamInterceptor(grpctrace.StreamClientInterceptor()),
	}

}
