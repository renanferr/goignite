package logger

import (
	"time"

	"github.com/b2wdigital/goignite/v2/core/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func streamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		logger := log.FromContext(context.Background())

		start := time.Now()
		wrapper := &recvWrapper{stream}
		err := handler(srv, wrapper)
		logger.Infof("invoke server method=%s duration=%s error=%v", info.FullMethod,
			time.Since(start), err)
		return err
	}
}

func unaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		logger := log.FromContext(ctx)

		start := time.Now()
		r, err := handler(ctx, req)
		logger.Infof("invoke server method=%s duration=%s error=%v response=%v", info.FullMethod,
			time.Since(start), err, r)
		return r, err
	}
}

type recvWrapper struct {
	grpc.ServerStream
}
