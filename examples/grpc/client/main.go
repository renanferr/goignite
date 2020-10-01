package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gigrpc "github.com/b2wdigital/goignite/grpc/v1/client"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	"github.com/prometheus/common/log"
)

func main() {

	ctx := context.Background()

	giconfig.Load()

	gilogrus.NewLogger()

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
		log.Fatalf("%v.Call(_) = _, %v", c, err)
	}

	gilog.Infof(test.Message)

	gilog.Infof(conn.GetState().String())
}
