package zap

import (
	"io"
	"os"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() log.Logger {

	cores := []zapcore.Core{}
	var writers []io.Writer

	if config.Bool(log.ConsoleEnabled) {
		level := getZapLevel(config.String(log.ConsoleLevel))
		writer := zapcore.Lock(os.Stdout)
		coreconsole := zapcore.NewCore(getEncoder(config.String(ConsoleFormatter)), writer, level)
		cores = append(cores, coreconsole)
		writers = append(writers, writer)
	}

	if config.Bool(log.FileEnabled) {
		s := []string{config.String(log.FilePath), "/", config.String(log.FileName)}
		fileLocation := strings.Join(s, "")

		lumber := &lumberjack.Logger{
			Filename: fileLocation,
			MaxSize:  config.Int(log.FileMaxSize),
			Compress: config.Bool(log.FileCompress),
			MaxAge:   config.Int(log.FileMaxAge),
		}

		level := getZapLevel(config.String(log.FileLevel))
		writer := zapcore.AddSync(lumber)
		corefile := zapcore.NewCore(getEncoder(config.String(FileFormatter)), writer, level)
		cores = append(cores, corefile)
		writers = append(writers, lumber)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	newlogger := &zapLogger{
		fields:        log.Fields{},
		sugaredLogger: logger,
		writers:       writers,
	}

	log.NewLogger(newlogger)
	return newlogger
}

func getEncoder(format string) zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	switch format {

	case "JSON":

		return zapcore.NewJSONEncoder(encoderConfig)

	default:

		return zapcore.NewConsoleEncoder(encoderConfig)

	}

}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "TRACE":
		return zapcore.DebugLevel
	case "WARN":
		return zapcore.WarnLevel
	case "DEBUG":
		return zapcore.DebugLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
	fields        log.Fields
	writers       []io.Writer
}

func (l *zapLogger) Tracef(format string, args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Trace(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *zapLogger) WithField(key string, value interface{}) log.Logger {

	fields := log.Fields{}
	fields[key] = value
	l.fields[key] = value

	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger, l.fields, l.writers}
}

func (l *zapLogger) Output() io.Writer {
	return io.MultiWriter(l.writers...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) WithFields(fields log.Fields) log.Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
		l.fields[k] = v
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger, l.fields, l.writers}
}

func (l *zapLogger) GetFields() log.Fields {
	return l.fields
}
