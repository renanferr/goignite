package logger

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/grpc/v1/client"
)

const (
	root    = client.ExtRoot + ".logger"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable logger")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
