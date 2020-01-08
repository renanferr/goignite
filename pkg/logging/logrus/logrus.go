package logrus

import (
	"io/ioutil"
	"os"
	"strings"

	c "github.com/jpfaria/goignite/pkg/config"
	. "github.com/jpfaria/goignite/pkg/logging/config"
	"github.com/jpfaria/goignite/pkg/logging/logrus/config"
	"github.com/jpfaria/logrus-redis-hook"
	"github.com/ravernkoh/cwlogsfmt"
	log "github.com/sirupsen/logrus"
)

func Start() {

	if c.Instance.Bool(config.RedisEnabled) {

		hookConfig := logredis.HookConfig{
			Host:   c.Instance.String(config.RedisHost),
			Key:    c.Instance.String(config.RedisKey),
			Format: c.Instance.String(config.RedisFormat),
			App:    c.Instance.String(config.RedisApp),
			Port:   c.Instance.Int(config.RedisPort),
			DB:     c.Instance.Int(config.RedisDb),
		}

		hook, err := logredis.NewHook(hookConfig)
		if err == nil {
			log.AddHook(hook)
		} else {
			log.Errorf("logredis error: %q", err)
		}

	}

	if c.Instance.Bool(LogFileEnabled) {

		formatter := getFormatter(c.Instance.String(config.FileFormatter))

		log.SetFormatter(formatter)

		s := []string{c.Instance.String(LogFilePath), "/", c.Instance.String(LogFileName)}
		logFile := strings.Join(s, "")

		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		log.SetOutput(f)

	} else if c.Instance.Bool(LogConsoleEnabled) {

		formatter := getFormatter(c.Instance.String(config.ConsoleFormatter))

		log.SetFormatter(formatter)

		log.SetOutput(os.Stdout)

	} else {

		log.SetOutput(ioutil.Discard)

	}

	switch level := c.Instance.String(LogLevel); level {

	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)

	}

}

func getFormatter(format string) log.Formatter {

	var formatter log.Formatter

	switch format {

	case "JSON":

		fmt := &log.JSONFormatter{
			FieldMap: log.FieldMap{
				log.FieldKeyTime:  "date",
				log.FieldKeyLevel: "log_level",
				log.FieldKeyMsg:   "log_message",
			},
		}

		fmt.TimestampFormat = c.Instance.String(c.FormatTimestamp)

		formatter = fmt

	case "AWS_CLOUD_WATCH":

		formatter = &cwlogsfmt.CloudWatchLogsFormatter{
			PrefixFields:     []string{"RequestId"},
			QuoteEmptyFields: true,
		}

	default:

		fmt := &log.TextFormatter{}
		fmt.TimestampFormat = c.Instance.String(c.FormatTimestamp)

		formatter = fmt

	}

	return formatter
}
