package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
	gimongo "github.com/b2wdigital/goignite/v2/mongo/v1"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	conn, err := gimongo.NewDefaultConn(context.Background())
	if err != nil {
		gilog.Panic(err)
	}

	err = conn.Client.Ping(context.Background(), nil)
	if err != nil {
		gilog.Panic(err)
	}

}
