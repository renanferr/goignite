package ginraws

import (
	giconfig "github.com/b2wdigital/goignite/config"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
)

const (
	ConfigRoot = ginewrelic.ConfigRoot + ".ext.aws"
	enabled    = ConfigRoot + ".enabled"
)

func init() {

	giconfig.Add(enabled, true, "enable/disable aws integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
