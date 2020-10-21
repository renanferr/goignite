package ginewrelic

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Logger struct {
}

func NewLogger() newrelic.Logger {
	return &Logger{}
}

func (l *Logger) Error(msg string, context map[string]interface{}) {
	gilog.WithFields(context).Error(msg)
}

func (l *Logger) Warn(msg string, context map[string]interface{}) {
	gilog.WithFields(context).Warn(msg)
}

func (l *Logger) Info(msg string, context map[string]interface{}) {
	gilog.WithFields(context).Info(msg)
}

func (l *Logger) Debug(msg string, context map[string]interface{}) {
	gilog.WithFields(context).Debug(msg)
}

func (l *Logger) DebugEnabled() bool {
	return giconfig.String(gilog.ConsoleLevel) == "DEBUG" || giconfig.String(gilog.FileLevel) == "DEBUG" ||
		giconfig.String(gilog.ConsoleLevel) == "TRACE" || giconfig.String(gilog.FileLevel) == "TRACE"
}
