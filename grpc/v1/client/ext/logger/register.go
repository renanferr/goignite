package logger

import (
	"google.golang.org/grpc"
)

func Register() []grpc.DialOption {

	if !IsEnabled() {
		return nil
	}

	return []grpc.DialOption{
		grpc.WithChainStreamInterceptor(streamInterceptor()),
		grpc.WithChainUnaryInterceptor(unaryInterceptor()),
	}
}
