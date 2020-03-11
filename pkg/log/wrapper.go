package log

// A global variable so that l functions can be directly accessed
var l Logger

// NewLogger returns an instance of logger
func NewLogger(logger Logger) {
	l = logger
}

func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return l.WithFields(keyValues)
}

func GetFields() Fields {
	return l.GetFields()
}
