package logger

import (
	"google.golang.org/grpc"
)

func Register() []grpc.ServerOption {

	if !IsEnabled() {
		return nil
	}

	return []grpc.ServerOption{
		grpc.ChainStreamInterceptor(streamInterceptor()),
		grpc.ChainUnaryInterceptor(unaryInterceptor()),
	}
}
