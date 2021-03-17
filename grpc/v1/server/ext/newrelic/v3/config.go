package newrelic

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/grpc/v1/server"
)

const (
	root    = server.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
