package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/config"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	gigocql "github.com/b2wdigital/goignite/gocql/v0"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	session, err := gigocql.NewDefaultSession(context.Background())
	defer session.Close()

	if err != nil {
		panic(err)
	}

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
