package giecho

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	echoRoot          = "gi.echo"
	hideBanner        = echoRoot + ".hidebanner"
	port              = echoRoot + ".port"
	statusRoute       = echoRoot + ".route.status"
	healthRoute       = echoRoot + ".route.health"
	jsonPrettyEnabled = echoRoot + ".json.pretty.enabled"
	extRoot           = echoRoot + ".ext"
	ConfigRoot        = extRoot + ".config"
	RouteRoot         = extRoot + ".route"
	MiddlewareRoot    = extRoot + ".middleware"
)

func init() {

	log.Println("getting configurations for echo")

	giconfig.Add(hideBanner, true, "echo hide/show banner")
	giconfig.Add(port, 8080, "server http port")
	giconfig.Add(healthRoute, "/health", "define health url")
	giconfig.Add(jsonPrettyEnabled, false, "enable/disable json pretty response")
}

func GetHideBanner() bool {
	return giconfig.Bool(hideBanner)
}

func GetPort() int {
	return giconfig.Int(port)
}

func GetStatusRoute() string {
	return giconfig.String(statusRoute)
}

func GetHealthRoute() string {
	return giconfig.String(healthRoute)
}

func GetJSONPrettyEnabled() bool {
	return giconfig.Bool(jsonPrettyEnabled)
}
