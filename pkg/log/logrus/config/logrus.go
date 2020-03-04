package config

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const RedisEnabled = "log.logrus.redis.enabled"
const RedisHost = "log.logrus.redis.host"
const RedisKey = "log.logrus.redis.key"
const RedisFormat = "log.logrus.redis.format"
const RedisApp = "log.logrus.redis.app"
const RedisPort = "log.logrus.redis.port"
const RedisDb = "log.logrus.redis.db"
const ConsoleFormatter = "log.logrus.console.formatter"
const FileFormatter = "log.logrus.file.formatter"

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
