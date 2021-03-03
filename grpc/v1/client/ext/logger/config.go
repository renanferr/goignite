package logger

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gigrpc "github.com/b2wdigital/goignite/v2/grpc/v1/client"
)

const (
	root    = gigrpc.ExtRoot + ".logger"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable logger")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
