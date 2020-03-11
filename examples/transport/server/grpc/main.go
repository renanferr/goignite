package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus/v1"
	"github.com/b2wdigital/goignite/pkg/transport/server/grpc/v1"
)

func main() {

	ctx := context.Background()

	config.Load()

	// start logrus
	log.NewLogger(logrus.NewLogger())

	s := grpc.Start(ctx)

	RegisterExampleServer(s, &Service{})

	grpc.Serve(ctx)
}

type Service struct {
}

func (h *Service) Test(ctx context.Context, request *TestRequest) (*TestResponse, error) {

	l := log.FromContext(ctx)

	l.Infof(request.Message)

	return &TestResponse{Message: "hello world"}, nil
}
