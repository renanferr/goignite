package logrus

import (
	"io/ioutil"
	"os"
	"strings"

	c "github.com/b2wdigital/goignite/pkg/config"
	. "github.com/b2wdigital/goignite/pkg/log/config"
	"github.com/b2wdigital/goignite/pkg/log/logrus/config"
	"github.com/b2wdigital/logrus-redis-hook"
	"github.com/ravernkoh/cwlogsfmt"
	log "github.com/sirupsen/logrus"
)

func Start() {

	if c.Bool(config.RedisEnabled) {

		hookConfig := logredis.HookConfig{
			Host:   c.String(config.RedisHost),
			Key:    c.String(config.RedisKey),
			Format: c.String(config.RedisFormat),
			App:    c.String(config.RedisApp),
			Port:   c.Int(config.RedisPort),
			DB:     c.Int(config.RedisDb),
		}

		hook, err := logredis.NewHook(hookConfig)
		if err == nil {
			log.AddHook(hook)
		} else {
			log.Errorf("logredis error: %q", err)
		}

	}

	if c.Bool(FileEnabled) {

		formatter := getFormatter(c.String(config.FileFormatter))

		log.SetFormatter(formatter)

		s := []string{c.String(FilePath), "/", c.String(FileName)}
		logFile := strings.Join(s, "")

		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		log.SetOutput(f)

	} else if c.Bool(ConsoleEnabled) {

		formatter := getFormatter(c.String(config.ConsoleFormatter))

		log.SetFormatter(formatter)

		log.SetOutput(os.Stdout)

	} else {

		log.SetOutput(ioutil.Discard)

	}

	switch level := c.String(Level); level {

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

		fmt.TimestampFormat = c.String(c.FormatTimestamp)

		formatter = fmt

	case "AWS_CLOUD_WATCH":

		formatter = &cwlogsfmt.CloudWatchLogsFormatter{
			PrefixFields:     []string{"RequestId"},
			QuoteEmptyFields: true,
		}

	default:

		fmt := &log.TextFormatter{}
		fmt.TimestampFormat = c.String(c.FormatTimestamp)

		formatter = fmt

	}

	return formatter
}
