package grpc

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func DebugStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		l := log.FromContext(context.Background())

		start := time.Now()
		wrapper := &recvWrapper{stream}
		err := handler(srv, wrapper)
		l.Debugf("invoke server method=%s duration=%s error=%v", info.FullMethod,
			time.Since(start), err)
		return err
	}
}

func DebugUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		l := log.FromContext(ctx)

		start := time.Now()
		r, err := handler(ctx, req)
		l.Debugf("invoke server method=%s duration=%s error=%v response=%v", info.FullMethod,
			time.Since(start), err, r)
		return r, err
	}
}

type recvWrapper struct {
	grpc.ServerStream
}
