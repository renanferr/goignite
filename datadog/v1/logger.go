package gidatadog

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

type Logger struct {
}

func NewLogger() ddtrace.Logger {
	return &Logger{}
}

func (l *Logger) Log(msg string) {

	var fn func(args ...interface{})

	switch giconfig.String(logLevel) {
	case "INFO":
		fn = gilog.Info
	case "DEBUG":
		fn = gilog.Debug
	default:
		fn = gilog.Debug
	}

	fn(msg)
}
