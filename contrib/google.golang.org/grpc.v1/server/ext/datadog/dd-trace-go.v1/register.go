package dd_trace_go_v1

import (
	"context"

	datadog "github.com/b2wdigital/goignite/v2/contrib/datadog/dd-trace-go.v1"
	"github.com/b2wdigital/goignite/v2/core/log"
	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("datadog interceptor successfully enabled in grpc server")

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpctrace.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(grpctrace.StreamServerInterceptor()),
	}

}
