package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const RedisEnabled = "log.redis.enabled"
const RedisHost = "log.redis.host"
const RedisKey = "log.redis.key"
const RedisFormat = "log.redis.format"
const RedisApp = "log.redis.app"
const RedisPort = "log.redis.port"
const RedisDb = "log.redis.db"
const ConsoleFormatter = "log.console.formatter"
const FileFormatter = "log.file.formatter"

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
