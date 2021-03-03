package logger

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gigrpc "github.com/b2wdigital/goignite/grpc/v1/client"
)

const (
	root    = gigrpc.ExtRoot + ".prometheus"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable prometheus")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
