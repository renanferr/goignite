package prometheus

import (
	"context"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() {
		return nil
	}

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.ChainStreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	}

}
