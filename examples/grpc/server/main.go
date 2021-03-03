package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gigrpc "github.com/b2wdigital/goignite/v2/grpc/v1/server"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
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

	logger := gilog.FromContext(ctx)

	logger.Infof(request.Message)

	return &TestResponse{Message: "hello world"}, nil
}
