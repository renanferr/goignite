package newrelic

import (
	"context"

	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() {
		return nil
	}

	app := ginewrelic.Application()

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(nrgrpc.UnaryServerInterceptor(app)),
		grpc.ChainStreamInterceptor(nrgrpc.StreamServerInterceptor(app)),
	}

}
