package giecho

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	echoRoot          = "gi.echo"
	hideBanner        = echoRoot + ".hidebanner"
	port              = echoRoot + ".port"
	jsonPrettyEnabled = echoRoot + ".json.pretty.enabled"
	extRoot           = echoRoot + ".ext"
	ConfigRoot        = extRoot + ".config"
	RouteRoot         = extRoot + ".route"
	MiddlewareRoot    = extRoot + ".middleware"
	IntegrateRoot     = extRoot + ".integrate"
)

func init() {
	giconfig.Add(hideBanner, true, "echo hide/show banner")
	giconfig.Add(port, 8080, "server http port")
	giconfig.Add(jsonPrettyEnabled, false, "enable/disable json pretty response")
}

func GetHideBanner() bool {
	return giconfig.Bool(hideBanner)
}

func GetPort() int {
	return giconfig.Int(port)
}

func GetJSONPrettyEnabled() bool {
	return giconfig.Bool(jsonPrettyEnabled)
}
