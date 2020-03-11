package main

import (
	"context"

	resty2 "github.com/b2wdigital/goignite/pkg/client/http/resty"
	"github.com/b2wdigital/goignite/pkg/client/http/resty/v2"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
)

func main() {

	var err error

	err = config.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	log.NewLogger(logrus.NewLogger())

	l := log.FromContext(ctx)

	client := resty.NewClient(ctx, &resty2.Options{})
	request := client.R().EnableTrace()

	response, err := request.Get("http://google.com")
	if err != nil {
		l.Fatalf(err.Error())
	}

	l.Infof(response.String())
}
