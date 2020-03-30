package chi

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	StatusRoute                = "transport.client.chi.route.status"
	HealthRoute                = "transport.client.chi.route.health"
	MiddlewareLogEnabled       = "transport.client.chi.middleware.log.enabled"
	MiddlewareRecoverEnabled   = "transport.client.chi.middleware.recover.enabled"
	MiddlewareRealIPEnabled    = "transport.client.chi.middleware.realip.enabled"
	MiddlewareRequestIDEnabled = "transport.client.chi.middleware.requestid.enabled"
	MiddlewareNewTID           = "transport.client.chi.middleware.newtid.enabled"
	MiddlewareNewRelic         = "transport.client.chi.middleware.newrelic.enabled"
)

func init() {

	log.Println("getting configurations for chi")

	config.Add(StatusRoute, "/resource-status", "define status url")
	config.Add(HealthRoute, "/health", "define health url")
	config.Add(MiddlewareLogEnabled, true, "enable/disable logging request middleware")
	config.Add(MiddlewareRecoverEnabled, true, "enable/disable recover middleware")
	config.Add(MiddlewareRealIPEnabled, true, "enable/disable real ip middleware")
	config.Add(MiddlewareRequestIDEnabled, true, "enable/disable request id middleware")
	config.Add(MiddlewareNewTID, true, "enable/disable new tid middleware")
	config.Add(MiddlewareNewRelic, false, "enable/disable new relic middleware")
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

func GetMiddlewareNewTidEnabled() bool {
	return config.Bool(MiddlewareNewTID)
}

func GetMiddlewareNewRelicEnabled() bool {
	return config.Bool(MiddlewareNewRelic)
}
