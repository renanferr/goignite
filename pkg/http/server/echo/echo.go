package echo

import (
	"strconv"

	srvconfig "github.com/jpfaria/goignite/pkg/http/server/config"
	"github.com/jpfaria/goignite/pkg/http/server/echo/config"
	"github.com/jpfaria/goignite/pkg/http/server/echo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	m "github.com/neko-neko/echo-logrus"
	"github.com/neko-neko/echo-logrus/log"
)

var (
	instance *echo.Echo
)

func Start() *echo.Echo {

	instance = echo.New()

	instance.HideBanner = config.GetHideBanner()
	instance.Logger = log.Logger()

	setDefaultMiddlewares(instance)
	setDefaultRouters(instance)

	return instance
}

func setDefaultMiddlewares(instance *echo.Echo) {
	instance.Use(m.Logger())
	instance.Use(middleware.Recover())
}

func setDefaultRouters(instance *echo.Echo) {
	statusRoute := srvconfig.GetStatusRoute()

	log.Infof("configuring status router on %s", statusRoute)

	statusHandler := handler.NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

	healthRoute := srvconfig.GetHealthRoute()

	log.Infof("configuring health router on %s", healthRoute)

	healthHandler := handler.NewHealthHandler()
	instance.GET(healthRoute, healthHandler.Get)
}

func Serve() {
	log.Info("starting echo server. https://echo.labstack.com/")
	instance.Logger.Fatal(instance.Start(getServerPort()))
}

func getServerPort() string {
	return ":" + strconv.Itoa(srvconfig.GetPort())
}
