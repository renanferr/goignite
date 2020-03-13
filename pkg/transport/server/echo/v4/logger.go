package echo

import (
	"fmt"
	"io"

	l "github.com/b2wdigital/goignite/pkg/log"
	"github.com/labstack/gommon/log"
)

type wrappedLogger struct {
	logger l.Logger
}

func (wl wrappedLogger) Output() io.Writer {
	wl.logger.Errorf("Output(): implement me")
	return nil
}

func (wl wrappedLogger) Prefix() string {
	wl.logger.Errorf("Prefix(): implement me")
	return ""
}

func (wl wrappedLogger) SetPrefix(p string) {
	wl.logger.Errorf("SetPrefix(p string): implement me")
}

func (wl wrappedLogger) Level() log.Lvl {
	wl.logger.Errorf("Level(): implement me")
	return log.INFO
}

func (wl wrappedLogger) SetLevel(v log.Lvl) {
	panic("SetLevel(v log.Lvl): implement me")
}

func (wl wrappedLogger) SetHeader(h string) {
	panic("SetHeader(h string): implement me")
}

func (wl wrappedLogger) Printj(j log.JSON) {
	panic("Printj(j log.JSON): implement me")
}

func (wl wrappedLogger) Debugj(j log.JSON) {
	panic("Debugj(j log.JSON) implement me")
}

func (wl wrappedLogger) Infoj(j log.JSON) {
	panic("Infoj(j log.JSON): implement me")
}

func (wl wrappedLogger) Warnj(j log.JSON) {
	panic("Warnj(j log.JSON): implement me")
}

func (wl wrappedLogger) Errorj(j log.JSON) {
	panic("Errorj(j log.JSON): implement me")
}

func (wl wrappedLogger) Fatalj(j log.JSON) {
	panic("Fatalj(j log.JSON): implement me")
}

func (wl wrappedLogger) Panic(i ...interface{}) {
	panic("Panic(i ...interface{}): implement me")
}

func (wl wrappedLogger) Panicj(j log.JSON) {
	panic("Panicj(j log.JSON): implement me")
}

func (wl wrappedLogger) Panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func (wl wrappedLogger) SetOutput(w io.Writer) {
	wl.logger.Panicf("Zap can only have its output set at creation time with the Output() option. SetOutput() does not work.")
}

func (wl wrappedLogger) Print(i ...interface{}) {
	wl.Info(i)
}

func (wl wrappedLogger) Printf(s string, i ...interface{}) {
	wl.Infof(s, i...)
}

func (wl wrappedLogger) Debug(i ...interface{}) {
	wl.logger.Debugf(fmt.Sprint(i...))
}

func (wl wrappedLogger) Debugf(s string, i ...interface{}) {
	wl.logger.Debugf(fmt.Sprintf(s, i...))
}

func (wl wrappedLogger) Info(i ...interface{}) {
	wl.logger.Infof(fmt.Sprint(i...))
}

func (wl wrappedLogger) Infof(s string, i ...interface{}) {
	wl.logger.Infof(fmt.Sprintf(s, i...))
}

func (wl wrappedLogger) Warn(i ...interface{}) {
	wl.logger.Warnf(fmt.Sprint(i...))
}

func (wl wrappedLogger) Warnf(s string, i ...interface{}) {
	wl.logger.Warnf(fmt.Sprintf(s, i...))
}

func (wl wrappedLogger) Error(i ...interface{}) {
	wl.logger.Errorf(fmt.Sprint(i...))
}

func (wl wrappedLogger) Errorf(s string, i ...interface{}) {
	wl.logger.Errorf(fmt.Sprintf(s, i...))
}

func (wl wrappedLogger) Fatal(i ...interface{}) {
	wl.logger.Fatalf(fmt.Sprint(i...))
}

func (wl wrappedLogger) Fatalf(s string, i ...interface{}) {
	wl.logger.Fatalf(fmt.Sprintf(s, i...))
}

func Wrap(l l.Logger) wrappedLogger {
	return wrappedLogger{logger: l}
}
