package main

import (
	"context"
	"log"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
	"github.com/b2wdigital/goignite/v2/mongo/v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

	conn, err := mongo.NewDefaultConn(context.Background())
	if err != nil {
		log.Panic(err)
	}

	err = conn.Client.Ping(context.Background(), nil)
	if err != nil {
		log.Panic(err)
	}

}
