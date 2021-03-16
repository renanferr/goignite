package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gigrpc "github.com/b2wdigital/goignite/v2/grpc/v1/client"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	ctx := context.Background()

	giconfig.Load()

	logger := gilogrus.NewLogger()

	request := &TestRequest{
		Message: "mensagem da requisição",
	}

	options := gigrpc.OptionsBuilder.
		Host("localhost").
		Port(9090).
		Build()

	conn := gigrpc.NewClientConn(ctx, &options)
	defer conn.Close()

	c := NewExampleClient(conn)

	test, err := c.Test(ctx, request)
	if err != nil {
		logger.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	gilog.Infof(test.Message)

	gilog.Infof(conn.GetState().String())
}
