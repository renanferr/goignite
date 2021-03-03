package logger

import (
	"google.golang.org/grpc"
)

func Register() []grpc.ServerOption {

	if !isEnabled() {
		return nil
	}

	return []grpc.ServerOption{
		grpc.StreamInterceptor(streamInterceptor()),
		grpc.UnaryInterceptor(unaryInterceptor()),
	}
}
