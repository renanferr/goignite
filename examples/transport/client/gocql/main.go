package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus/v1"
	"github.com/b2wdigital/goignite/pkg/transport/client/gocql/v0"
)

func main() {

	config.Load()

	logrus.NewLogger()

	session, err := gocql.NewDefaultSession(context.Background())
	defer session.Close()

	if err != nil {
		panic(err)
	}

	all := health.CheckAll(context.Background())

	log.Info(all)

}
