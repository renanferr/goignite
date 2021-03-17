package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/grpc/v1/client"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	ctx := context.Background()

	config.Load()

	logger := logrus.NewLogger()

	request := &TestRequest{
		Message: "mensagem da requisição",
	}

	options := client.OptionsBuilder.
		Host("localhost").
		Port(9090).
		Build()

	conn := client.NewClientConn(ctx, &options)
	defer conn.Close()

	c := NewExampleClient(conn)

	test, err := c.Test(ctx, request)
	if err != nil {
		logger.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	log.Infof(test.Message)

	log.Infof(conn.GetState().String())
}
