package logger

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gigrpc "github.com/b2wdigital/goignite/grpc/v1/client"
)

const (
	root    = gigrpc.ExtRoot + ".logger"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable logger")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
