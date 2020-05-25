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

	// call google

	googleopt := new(giresty.Options)

	err = giconfig.UnmarshalWithPath("app.client.resty.google", googleopt)
	if err != nil {
		l.Errorf(err.Error())
	}

	cligoogle := giresty.NewClient(ctx, googleopt)
	reqgoogle := cligoogle.R()

	respgoogle, err := reqgoogle.Get("/")
	if err != nil {
		l.Fatalf(err.Error())
	}

	l.Infof(respgoogle.String())

	// call acom

	acomopt := new(giresty.Options)

	err = giconfig.UnmarshalWithPath("app.client.resty.acom", acomopt)
	if err != nil {
		l.Errorf(err.Error())
	}

	cliacom := giresty.NewClient(ctx, acomopt)
	reqacom := cliacom.R()

	respacom, err := reqacom.Get("/")
	if err != nil {
		l.Fatalf(err.Error())
	}

	l.Infof(respacom.String())
}
