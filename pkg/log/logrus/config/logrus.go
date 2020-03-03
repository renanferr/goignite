package config

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const RedisEnabled = "logging.logrus.redis.enabled"
const RedisHost = "logging.logrus.redis.host"
const RedisKey = "logging.logrus.redis.key"
const RedisFormat = "logging.logrus.redis.format"
const RedisApp = "logging.logrus.redis.app"
const RedisPort = "logging.logrus.redis.port"
const RedisDb = "logging.logrus.redis.db"
const ConsoleFormatter = "logging.logrus.console.formatter"
const FileFormatter = "logging.logrus.file.formatter"

func init() {

	log.Println("getting configurations for logrus")

	config.Add(RedisEnabled, false, "enable/disable redis logging")
	config.Add(RedisHost, "localhost", "redis host")
	config.Add(RedisKey, "", "redis key")
	config.Add(RedisFormat, "b2w", "redis format")
	config.Add(RedisApp, "logger", "redis app")
	config.Add(RedisPort, 6379, "redis port")
	config.Add(RedisDb, 0, "redis db")
	config.Add(ConsoleFormatter, "TEXT", "console formatter TEXT/JSON/AWS_CLOUD_WATCH")
	config.Add(FileFormatter, "JSON", "file formatter TEXT/JSON/AWS_CLOUD_WATCH")

}
