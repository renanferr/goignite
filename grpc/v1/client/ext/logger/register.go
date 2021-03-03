package logger

import (
	"google.golang.org/grpc"
)

func Register() []grpc.DialOption {

	if !isEnabled() {
		return nil
	}

	return []grpc.DialOption{
		grpc.WithStreamInterceptor(streamInterceptor()),
		grpc.WithUnaryInterceptor(unaryInterceptor()),
	}
}
