package zerolog

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() log.Logger {
	fileEnabled := config.Bool(log.FileEnabled)
	consoleEnabled := config.Bool(log.ConsoleEnabled)
	fields := make(log.Fields)

	format := config.String(Formatter)
	writer := getWriter(format, fileEnabled, consoleEnabled)
	if writer == nil {
		zerologger := zerolog.Nop()
		return &logger{
			logger: &zerologger,
			fields: fields,
		}
	}

	zerolog.MessageFieldName = "log_message"
	zerolog.LevelFieldName = "log_level"

	zerologger := zerolog.New(writer).With().Timestamp().Logger()
	level := getLogLevel(config.String(log.ConsoleLevel))
	zerologger = zerologger.Level(level)

	logger := &logger{
		logger: &zerologger,
		fields: fields,
		writer: writer,
	}

	log.NewLogger(logger)
	return logger
}

type logger struct {
	logger *zerolog.Logger
	fields log.Fields
	writer io.Writer
}

func getLogLevel(level string) zerolog.Level {
	switch level {
	case "DEBUG":
		return zerolog.DebugLevel
	case "WARN":
		return zerolog.WarnLevel
	case "FATAL":
		return zerolog.FatalLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "TRACE":
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}

func getWriter(format string, fileEnabled bool, consoleEnabled bool) io.Writer {
	var writer io.Writer
	switch format {
	case "TEXT":
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	default:
		writer = os.Stdout
	}

	if fileEnabled {
		s := []string{config.String(log.FilePath), "/", config.String(log.FileName)}
		fileLocation := strings.Join(s, "")

		fileHandler := &lumberjack.Logger{
			Filename: fileLocation,
			MaxSize:  config.Int(log.FileMaxSize),
			Compress: config.Bool(log.FileCompress),
			MaxAge:   config.Int(log.FileMaxAge),
		}

		if consoleEnabled {
			return io.MultiWriter(writer, fileHandler)
		} else {
			return fileHandler
		}
	} else if consoleEnabled {
		return writer
	}

	return nil
}

func (l *logger) Tracef(format string, args ...interface{}) {
	l.logger.Trace().Msgf(format, args...)
}

func (l *logger) Trace(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Trace().Msgf(format.String(), args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Debug().Msgf(format.String(), args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Info().Msgf(format.String(), args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *logger) Warn(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Warn().Msgf(format.String(), args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Error().Msgf(format.String(), args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Fatal().Msgf(format.String(), args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.logger.Panic().Msgf(format, args...)
}

func (l *logger) Panic(args ...interface{}) {
	format := bytes.NewBufferString("")
	for _ = range args {
		format.WriteString("%v")
	}

	l.logger.Panic().Msgf(format.String(), args...)
}

func (l *logger) WithField(key string, value interface{}) log.Logger {

	newFields := log.Fields{}
	newFields[key] = value

	newLogger := l.logger.With().Fields(newFields).Logger()
	return &logger{&newLogger, newFields, l.writer}
}

func (l *logger) WithFields(fields log.Fields) log.Logger {
	newFields := log.Fields{}

	for k, v := range l.fields {
		newFields[k] = v
	}

	for k, v := range fields {
		newFields[k] = v
	}

	l.eraseFields()
	newLogger := l.logger.With().Fields(newFields).Logger()
	return &logger{&newLogger, newFields, l.writer}
}

func (l *logger) GetFields() log.Fields {
	return l.fields
}

func (l *logger) Output() io.Writer {
	return l.writer
}

func (l *logger) eraseFields() {
	// zerolog does no de-duplication of fields
	// we cover this generating a new zerolog.Context
	l.logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return zerolog.Context{}
	})
}
