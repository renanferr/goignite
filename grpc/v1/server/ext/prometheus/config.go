package prometheus

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gigrpc "github.com/b2wdigital/goignite/grpc/v1/server"
)

const (
	root    = gigrpc.ExtRoot + ".prometheus"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable prometheus")
	giconfig.Add(route, "/metrics", "prometheus route")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetRoute() string {
	return giconfig.String(enabled)
}
