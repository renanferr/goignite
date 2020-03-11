package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/server/grpc/v2"
)

func main() {

	ctx := context.Background()

	var err error

	err = config.Load()
	if err != nil {
		panic(err)
	}

	// start logrus
	log.NewLogger(v1.NewLogger())

	s := v1.Start(ctx)

	RegisterExampleServer(s, &Service{})

	v1.Serve(ctx)
}

type Service struct {
}

func (h *Service) Test(ctx context.Context, request *TestRequest) (*TestResponse, error) {

	l := log.FromContext(ctx)

	l.Infof(request.Message)

	return &TestResponse{Message: "hello world"}, nil
}
