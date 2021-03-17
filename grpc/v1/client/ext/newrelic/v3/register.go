package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.DialOption {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("newrelic interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
		grpc.WithChainStreamInterceptor(nrgrpc.StreamClientInterceptor),
	}

}
