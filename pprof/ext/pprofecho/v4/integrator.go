package gipprofecho

import (
	giecho "github.com/b2wdigital/goignite/echo/v4"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	echopprof "github.com/hiko1129/echo-pprof"
	"github.com/labstack/echo/v4"
)

type Integrator struct {
}

func Integrate() error {
	if !IsEnabled() {
		return nil
	}

	integrator := &Integrator{}
	return gieventbus.SubscribeOnce(giecho.TopicInstance, integrator.Integrate)
}

func (i *Integrator) Integrate(instance *echo.Echo) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating echo with pprof")

	echopprof.Wrap(instance)

	logger.Debug("pprof integrated with echo with success")

	return nil
}
