package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	giresty "github.com/b2wdigital/goignite/resty/v2"
	"github.com/b2wdigital/goignite/resty/v2/ext/health"
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

	client := giresty.NewClient(ctx, &giresty.Options{}, healthIntegrator.Integrate)
	request := client.R().EnableTrace()

	response, err := request.Get("http://google.com")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	logger.Infof(response.String())
}
