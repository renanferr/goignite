package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/client/grpc"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
)

func main() {

	ctx := context.Background()

	var err error

	err = config.Load()
	if err != nil {
		panic(err)
	}

	log.NewLogger(logrus.NewLogger())

	request := &TestRequest{
		Message: "mensagem da requisição",
	}

	health := grpc.OptionsHealthBuilder.
		Description("example grpc server").
		Required(true).
		Enabled(true).
		Build()

	options := grpc.OptionsBuilder.
		Host("localhost").
		Port(9090).
		Health(health).
		Build()

	conn := grpc.NewClient(ctx, &options)
	defer conn.Close()

	c := NewExampleClient(conn)

	test, err := c.Test(ctx, request)
	if err != nil {
		log.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	log.Infof(test.Message)

	log.Infof(conn.GetState().String())
}
