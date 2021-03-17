package newrelic

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/redis/v7"
)

const (
	root    = redis.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable redis integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
