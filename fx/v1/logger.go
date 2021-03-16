package gifx

import (
	gilog "github.com/b2wdigital/goignite/v2/log"
	"go.uber.org/fx"
)

type Printer struct {
}

func (*Printer) Printf(format string, args ...interface{}) {
	switch LogLevel() {
	case "INFO":
		gilog.Infof(format, args...)
	case "TRACE":
		gilog.Tracef(format, args...)
	default:
		gilog.Debugf(format, args...)
	}
}

func NewLogger() fx.Printer {
	return &Printer{}
}
