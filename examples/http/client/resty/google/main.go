package main

import (
	"context"
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/http/client/resty/model"
	"github.com/b2wdigital/goignite/pkg/http/client/resty/v2"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
)

func main() {

	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	logrus.Start()

	l := logrus.FromContext(ctx)

	client := resty.NewClient(ctx, &model.Options{})
	request := client.R().EnableTrace()

	response, err := request.Get("http://google.com")
	if err!= nil {
		l.Fatal(err)
	}

	l.Info(response.Request.Header)
	l.Info(response)
}
