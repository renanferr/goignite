package fx

import (
	"github.com/b2wdigital/goignite/v2/log"
	"go.uber.org/fx"
)

type Printer struct {
}

func (*Printer) Printf(format string, args ...interface{}) {
	switch LogLevel() {
	case "INFO":
		log.Infof(format, args...)
	case "TRACE":
		log.Tracef(format, args...)
	default:
		log.Debugf(format, args...)
	}
}

func NewLogger() fx.Printer {
	return &Printer{}
}
