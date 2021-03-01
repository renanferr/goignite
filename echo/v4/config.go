package giecho

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root              = "gi.echo"
	hideBanner        = root + ".hidebanner"
	port              = root + ".port"
	jsonPrettyEnabled = root + ".json.pretty.enabled"
	ExtRoot           = root + ".ext"
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
