package giawsnewrelic

import (
	giaws "github.com/b2wdigital/goignite/v2/aws/v2"
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root    = giaws.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
