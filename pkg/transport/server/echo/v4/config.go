package echo

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	HideBanner               = "transport.server.echo.hidebanner"
	Port                     = "transport.server.echo.port"
	StatusRoute              = "transport.server.echo.route.status"
	HealthRoute              = "transport.server.echo.route.health"
	MiddlewareLogEnabled     = "transport.server.echo.middleware.log.enabled"
	MiddlewareRecoverEnabled = "transport.server.echo.middleware.recover.enabled"
)

func init() {

	log.Println("getting configurations for echo")

	config.Add(HideBanner, true, "echo hide/show banner")
	config.Add(Port, 8080, "server http port")
	config.Add(StatusRoute, "/resource-status", "define status url")
	config.Add(HealthRoute, "/health", "define health url")
	config.Add(MiddlewareLogEnabled, true, "enable/disable logging request middleware")
	config.Add(MiddlewareRecoverEnabled, true, "enable/disable logging request middleware")

}

func GetHideBanner() bool {
	return config.Bool(HideBanner)
}

func GetPort() int {
	return config.Int(Port)
}

func GetStatusRoute() string {
	return config.String(StatusRoute)
}

func GetHealthRoute() string {
	return config.String(HealthRoute)
}

func GetMiddlewareLogEnabled() bool {
	return config.Bool(MiddlewareLogEnabled)
}

func GetMiddlewareRecoverEnabled() bool {
	return config.Bool(MiddlewareRecoverEnabled)
}
