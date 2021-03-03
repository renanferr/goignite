package semaphore

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	semaphoreRoot = giecho.ExtRoot + ".semaphore"
	enabled       = semaphoreRoot + ".enabled"
	limit         = semaphoreRoot + ".limit"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable semaphore middleware")
	giconfig.Add(limit, 10000, "defines numbers for concurrent connections")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetLimit() int {
	return giconfig.Int(limit)
}
