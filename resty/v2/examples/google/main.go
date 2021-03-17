package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
	giresty "github.com/b2wdigital/goignite/v2/resty/v2"
	health "github.com/b2wdigital/goignite/v2/resty/v2/ext/health"
	"github.com/go-resty/resty/v2"
)

func main() {

	var err error

	giconfig.Load()

	ctx := context.Background()

	gilogrus.NewLogger()

	logger := gilog.FromContext(ctx)

	options := health.OptionsBuilder.
		Host("http://google.com").
		Endpoint("/status").
		Name("Google Inc").
		Description("Search Engine").
		Required(true).
		Enabled(true).
		Build()

	healthIntegrator := health.NewIntegrator(&options)

	client := giresty.NewClient(ctx, &giresty.Options{}, healthIntegrator.Register)
	request := client.R().EnableTrace()

	var resp *resty.Response
	resp, err = request.Get("http://google.com")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if resp != nil {
		logger.Infof(resp.String())
	}
}
