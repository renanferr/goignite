package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/gocql/v0"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

	session, err := gocql.NewDefaultSession(context.Background())
	if err != nil {
		panic(err)
	}

	defer session.Close()

	err = session.Query("void").Exec()
	if err != nil {
		panic(err)
	}

}
