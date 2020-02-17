package main

import (
	"context"
	"log"

	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/http/client/resty/model"
	resty "github.com/jpfaria/goignite/pkg/http/client/resty/v2"
	"github.com/jpfaria/goignite/pkg/logging/logrus"
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
