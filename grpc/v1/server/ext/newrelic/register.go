package logger

import (
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

func Register() []grpc.ServerOption {

	if !isEnabled() {
		return nil
	}

	app := ginewrelic.Application()

	return []grpc.ServerOption{
		grpc.UnaryInterceptor(nrgrpc.UnaryServerInterceptor(app)),
		grpc.StreamInterceptor(nrgrpc.StreamServerInterceptor(app)),
	}

}
