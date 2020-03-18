package main

import (
	"context"
	"encoding/json"

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

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
