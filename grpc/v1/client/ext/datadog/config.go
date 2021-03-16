package gigrpcdatadog

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gigrpc "github.com/b2wdigital/goignite/v2/grpc/v1/client"
)

const (
	root    = gigrpc.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
