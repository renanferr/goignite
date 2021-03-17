package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gigocql "github.com/b2wdigital/goignite/v2/gocql/v0"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	session, err := gigocql.NewDefaultSession(context.Background())
	if err != nil {
		panic(err)
	}

	defer session.Close()

	err = session.Query("void").Exec()
	if err != nil {
		panic(err)
	}

}
