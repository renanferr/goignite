package logrus

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	logredis "github.com/jpfaria/logrus-redis-hook"
	"github.com/ravernkoh/cwlogsfmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() log.Logger {

	lLogger := new(logrus.Logger)

	if config.Bool(RedisEnabled) {

		hookConfig := logredis.HookConfig{
			Host:   config.String(RedisHost),
			Key:    config.String(RedisKey),
			Format: config.String(RedisFormat),
			App:    config.String(RedisApp),
			Port:   config.Int(RedisPort),
			DB:     config.Int(RedisDb),
		}

		hook, err := logredis.NewHook(hookConfig)
		if err == nil {
			lLogger.AddHook(hook)
		} else {
			lLogger.Errorf("logredis error: %q", err)
		}

	}

	var fileHandler *lumberjack.Logger

	lLogger.SetOutput(ioutil.Discard)

	if config.Bool(log.FileEnabled) {

		s := []string{config.String(log.FilePath), "/", config.String(log.FileName)}
		fileLocation := strings.Join(s, "")

		fileHandler = &lumberjack.Logger{
			Filename: fileLocation,
			MaxSize:  config.Int(log.FileMaxSize),
			Compress: config.Bool(log.FileCompress),
			MaxAge:   config.Int(log.FileMaxAge),
		}

	}

	if config.Bool(log.ConsoleEnabled) && config.Bool(log.FileEnabled) {
		lLogger.SetOutput(io.MultiWriter(os.Stdout, fileHandler))
	} else if config.Bool(log.FileEnabled) {
		lLogger.SetOutput(fileHandler)
	} else if config.Bool(log.ConsoleEnabled) {
		lLogger.SetOutput(os.Stdout)
	}

	level := getLogLevel(config.String(log.ConsoleLevel))
	formatter := getFormatter(config.String(Formatter))

	lLogger.SetLevel(level)
	lLogger.SetFormatter(formatter)

	return &logger{
		logger: lLogger,
	}

}

func getLogLevel(level string) logrus.Level {

	switch level {

	case "DEBUG":
		return logrus.DebugLevel
	case "WARN":
		return logrus.WarnLevel
	case "FATAL":
		return logrus.FatalLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "TRACE":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}

}

func getFormatter(format string) logrus.Formatter {

	var formatter logrus.Formatter

	switch format {

	case "JSON":

		fmt := &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "date",
				logrus.FieldKeyLevel: "log_level",
				logrus.FieldKeyMsg:   "log_message",
			},
		}

		fmt.TimestampFormat = config.String(config.FormatTimestamp)

		formatter = fmt

	case "AWS_CLOUD_WATCH":

		formatter = &cwlogsfmt.CloudWatchLogsFormatter{
			PrefixFields:     []string{"RequestId"},
			QuoteEmptyFields: true,
		}

	default:

		fmt := &logrus.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
		}
		fmt.TimestampFormat = config.String(config.FormatTimestamp)

		formatter = fmt

	}

	return formatter
}

type logger struct {
	logger *logrus.Logger
	fields log.Fields
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logger) WithFields(fields log.Fields) log.Logger {
	l.fields = fields
	return &logEntry{
		entry:  l.logger.WithFields(convertToLogrusFields(fields)),
		fields: fields,
	}
}

func (l *logger) GetFields() log.Fields {
	return l.fields
}

type logEntry struct {
	entry  *logrus.Entry
	fields log.Fields
}

func (l *logEntry) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logEntry) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logEntry) Panicf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logEntry) WithFields(fields log.Fields) log.Logger {
	return &logEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logEntry) GetFields() log.Fields {
	return l.fields
}

func convertToLogrusFields(fields log.Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
