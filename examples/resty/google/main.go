package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	giresty "github.com/b2wdigital/goignite/resty/v2"
)

func main() {

	var err error

	giconfig.Load()

	ctx := context.Background()

	gilogrus.NewLogger()

	l := gilog.FromContext(ctx)

	client := giresty.NewClient(ctx, &giresty.Options{})
	request := client.R().EnableTrace()

	response, err := request.Get("http://google.com")
	if err != nil {
		l.Fatalf(err.Error())
	}

	l.Infof(response.String())
}
