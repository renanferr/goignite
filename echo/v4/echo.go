package giecho

import (
	"context"
	"strconv"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/labstack/echo/v4"
)

var (
	instance *echo.Echo
)

func Start(ctx context.Context, exts ...func(context.Context, *echo.Echo) error) *echo.Echo {

	instance = echo.New()

	instance.HideBanner = GetHideBanner()
	instance.Logger = WrapLogger(gilog.GetLogger())

	for _, ext := range exts {
		if err := ext(ctx, instance); err != nil {
			panic(err)
		}
	}

	return instance
}

func Serve(ctx context.Context) {
	logger := gilog.FromContext(ctx)
	logger.Infof("starting echo server. https://echo.labstack.com/")
	err := instance.Start(serverPort())
	if err != nil {
		instance.Logger.Fatalf(err.Error())
	}
}

func serverPort() string {
	return ":" + strconv.Itoa(GetPort())
}
