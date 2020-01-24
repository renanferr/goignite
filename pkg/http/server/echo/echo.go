package echo

import (
	"strconv"

	sconfig "github.com/jpfaria/goignite/pkg/http/server/config"
	"github.com/jpfaria/goignite/pkg/http/server/echo/config"
	"github.com/jpfaria/goignite/pkg/http/server/echo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	m "github.com/neko-neko/echo-logrus"
	"github.com/neko-neko/echo-logrus/log"
	l "github.com/sirupsen/logrus"
)

var (
	instance *echo.Echo
)

func Start() *echo.Echo {

	instance = echo.New()

	instance.HideBanner = config.GetHideBanner()
	instance.Logger = log.Logger()

	setMiddlewares(instance)
	setDefaultRouters(instance)

	return instance
}

func setMiddlewares(instance *echo.Echo) {
	// Echo.Logger = logrusmiddleware.Logger{Logger: logrus.StandardLogger()}
	// Echo.Use(logrusmiddleware.Hook())

	// Echo.Use(middleware.RequestID())
	// Echo.Use(middleware.CORS())
	// Echo.Use(middleware.Gzip())
	instance.Use(m.Logger())
	instance.Use(middleware.Recover())
}

func setDefaultRouters(instance *echo.Echo) {
	statusRoute := sconfig.GetStatusRoute()
	log.Infof("configuring status router on %s", statusRoute)
	statusHandler := handler.NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)
}

func Serve() {
	l.Println("starting echo server")
	instance.Logger.Fatal(instance.Start(getServerPort()))
}

func getServerPort() string {
	return ":" + strconv.Itoa(sconfig.GetPort())
}
