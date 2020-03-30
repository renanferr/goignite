package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus/v1"
	"github.com/b2wdigital/goignite/pkg/transport/client/resty/v2"
)

func main() {

	var err error

	config.Load()

	ctx := context.Background()

	log.NewLogger(logrus.NewLogger())

	l := log.FromContext(ctx)

	// call google

	googleopt := new(resty.Options)

	err = config.UnmarshalWithPath("app.client.resty.google", googleopt)
	if err != nil {
		l.Errorf(err.Error())
	}

	cligoogle := resty.NewClient(ctx, googleopt)
	reqgoogle := cligoogle.R()

	respgoogle, err := reqgoogle.Get("/")
	if err != nil {
		l.Fatalf(err.Error())
	}

	l.Infof(respgoogle.String())

	// call acom

	acomopt := new(resty.Options)

	err = config.UnmarshalWithPath("app.client.resty.acom", acomopt)
	if err != nil {
		l.Errorf(err.Error())
	}

	cliacom := resty.NewClient(ctx, acomopt)
	reqacom := cliacom.R()

	respacom, err := reqacom.Get("/")
	if err != nil {
		l.Fatalf(err.Error())
	}

	l.Infof(respacom.String())
}
