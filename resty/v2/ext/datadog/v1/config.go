package girestydatadog

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	girest "github.com/b2wdigital/goignite/v2/resty/v2"
)

const (
	ConfigRoot = girest.ExtRoot + ".datadog"
	enabled    = ConfigRoot + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable datadog integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
