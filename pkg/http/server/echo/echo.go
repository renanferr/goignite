package echo

import (
	"strconv"

	c "github.com/jpfaria/goignite/pkg/config"
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

	instance.HideBanner = c.Instance.Bool(config.HideBanner)

	instance.Logger = log.Logger()

	// Echo.Logger = logrusmiddleware.Logger{Logger: logrus.StandardLogger()}
	// Echo.Use(logrusmiddleware.Hook())

	// Echo.Use(middleware.RequestID())
	// Echo.Use(middleware.CORS())
	// Echo.Use(middleware.Gzip())
	instance.Use(m.Logger())
	instance.Use(middleware.Recover())

	log.Infof("configuring status router on %s", c.Instance.String(sconfig.StatusRoute))

	statusHandler := handler.NewResourceStatusHandler()

	instance.GET(c.Instance.String(sconfig.StatusRoute), statusHandler.Get)

	return instance
}

func Serve() {
	l.Println("starting echo server")
	instance.Logger.Fatal(instance.Start(getServerPort()))
}

func getServerPort() string {
	return ":" + strconv.Itoa(c.Instance.Int(sconfig.Port))
}
