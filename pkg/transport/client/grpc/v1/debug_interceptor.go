package v1

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func DebugStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

		l := log.FromContext(ctx)

		start := time.Now()
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		l.Debugf("invoke server method=%s duration=%s error=%v", method,
			time.Since(start), err)
		return clientStream, err
	}
}

func DebugUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		l := log.FromContext(ctx)

		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		l.Debugf("invoke server method=%s duration=%s error=%v request=%v response=%v", method,
			time.Since(start), err, req, reply)
		return err
	}
}
