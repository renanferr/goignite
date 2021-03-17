package newrelic

import (
	ants "github.com/b2wdigital/goignite/v2/ants/v2"
	"github.com/b2wdigital/goignite/v2/config"
)

const (
	root    = ants.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
