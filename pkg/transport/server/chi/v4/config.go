package chi

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	StatusRoute              = "transport.server.chi.route.status"
	HealthRoute              = "transport.server.chi.route.health"
	MiddlewareLogEnabled     = "transport.server.chi.middleware.log.enabled"
	MiddlewareRecoverEnabled = "transport.server.chi.middleware.recover.enabled"
	MiddlewareRealIPEnabled = "transport.server.chi.middleware.RealIP.enabled"
	MiddlewareRequestIDEnabled = "transport.server.chi.middleware.RequestID.enabled"
)

func init() {

	log.Println("getting configurations for chi")

	config.Add(StatusRoute, "/resource-status", "define status url")
	config.Add(HealthRoute, "/health", "define health url")
	config.Add(MiddlewareLogEnabled, true, "enable/disable logging request middleware")
	config.Add(MiddlewareRecoverEnabled, true, "enable/disable recover middleware")
	config.Add(MiddlewareRealIPEnabled, true, "enable/disable real ip middleware")
	config.Add(MiddlewareRequestIDEnabled, true, "enable/disable request id middleware")
}

func GetStatusRoute() string {
	return config.String(StatusRoute)
}

func GetHealthRoute() string {
	return config.String(HealthRoute)
}

func GetMiddlewareRecoverEnabled() bool {
	return config.Bool(MiddlewareRecoverEnabled)
}
func GetMiddlewareLoggerEnabled() bool {
	return config.Bool(MiddlewareLogEnabled)
}
func GetMiddlewareRealIPEnabled() bool {
	return config.Bool(MiddlewareRealIPEnabled)
}
func GetMiddlewareRequestIDEnabled() bool {
	return config.Bool(MiddlewareRequestIDEnabled)
}


