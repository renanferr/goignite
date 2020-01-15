package main

import (
	"context"

	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/http/client/resty/model"
	resty "github.com/jpfaria/goignite/pkg/http/client/resty/v2"
	"github.com/jpfaria/goignite/pkg/logging/logrus"
)

func main() {

	config.Parse()
	log := logrus.FromContext(context.Background())

	client := resty.NewClient(&model.Options{})
	request := client.R().EnableTrace()



	response, err := request.Get("http://google.com")
	if err!= nil {
		log.Fatal(err)
	}

	log.Info(response.Request.Header)
	log.Info(response)

}
