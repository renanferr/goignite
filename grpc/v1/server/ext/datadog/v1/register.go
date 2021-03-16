package gigrpcdatadog

import (
	"context"

	gidatadog "github.com/b2wdigital/goignite/v2/datadog/v1"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() || !gidatadog.IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)
	logger.Debug("datadog interceptor successfully enabled in grpc server")

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpctrace.UnaryServerInterceptor()),
		grpc.ChainStreamInterceptor(grpctrace.StreamServerInterceptor()),
	}

}
