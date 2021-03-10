package giecho

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
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
	giconfig.Add(port, 8080, "Server http port")
	giconfig.Add(jsonPrettyEnabled, false, "enable/disable json pretty response")
}

func GetJSONPrettyEnabled() bool {
	return giconfig.Bool(jsonPrettyEnabled)
}
