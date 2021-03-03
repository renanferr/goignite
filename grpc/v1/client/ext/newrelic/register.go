package newrelic

import (
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

func Register() []grpc.DialOption {

	if !IsEnabled() {
		return nil
	}

	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
		grpc.WithChainStreamInterceptor(nrgrpc.StreamClientInterceptor),
	}

}
