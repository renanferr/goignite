package prometheus

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func Register() []grpc.ServerOption {

	if !IsEnabled() {
		return nil
	}

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.ChainStreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	}

}
