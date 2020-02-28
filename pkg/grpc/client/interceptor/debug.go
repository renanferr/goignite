package interceptor

import (
	"time"

	"github.com/jpfaria/goignite/pkg/log/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func DebugStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

		log := logrus.FromContext(ctx)

		start := time.Now()
		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		log.Debugf("invoke server method=%s duration=%s error=%v", method,
			time.Since(start), err)
		return clientStream, err
	}
}

func DebugUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		log := logrus.FromContext(ctx)

		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Debugf("invoke server method=%s duration=%s error=%v request=%v response=%v", method,
			time.Since(start), err, req, reply)
		return err
	}
}
