package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/b2wdigital/goignite/pkg/server/grpc"
)

func main() {

	ctx := context.Background()

	var err error

	err = config.Load()
	if err != nil {
		panic(err)
	}

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
