package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/grpc/client"
	"github.com/b2wdigital/goignite/pkg/grpc/client/model"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
)

func main() {

	ctx := context.Background()

	config.Parse()

	logrus.Start()

	log := logrus.FromContext(ctx)

	request := &TestRequest{
		Message: "mensagem da requisição",
	}

	health := model.OptionsHealthBuilder.
		Description("example grpc server").
		Required(true).
		Enabled(true).
		Build()

	options := model.OptionsBuilder.
		Host("localhost").
		Port(9090).
		Health(health).
		Build()

	conn := client.NewClient(ctx, &options)
	defer conn.Close()

	c := NewExampleClient(conn)

	test, err := c.Test(ctx, request)
	if err != nil {
		log.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	log.Info(test.Message)

	log.Info(conn.GetState())
}
