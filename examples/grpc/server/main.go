package main

import (
	"context"
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/grpc/server"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
)

func main() {

	ctx := context.Background()

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logrus.Start()

	s := server.Start(ctx)

	RegisterExampleServer(s, &Service{})

	server.Serve(ctx)
}


type Service struct {
}

func (h *Service) Test(ctx context.Context, request *TestRequest) (*TestResponse, error) {

	l := logrus.FromContext(ctx)

	l.Info(request.Message)

	return &TestResponse{Message: "hello world"}, nil
}
