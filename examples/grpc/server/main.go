package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gigrpc "github.com/b2wdigital/goignite/grpc/v1/server"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
)

func main() {

	ctx := context.Background()

	giconfig.Load()

	// start logrus
	gilogrus.NewLogger()

	s := gigrpc.Start(ctx)

	RegisterExampleServer(s, &Service{})

	gigrpc.Serve(ctx)
}

type Service struct {
}

func (h *Service) Test(ctx context.Context, request *TestRequest) (*TestResponse, error) {

	l := gilog.FromContext(ctx)

	l.Infof(request.Message)

	return &TestResponse{Message: "hello world"}, nil
}
