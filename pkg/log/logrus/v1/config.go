package logrus

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	RedisEnabled = "log.logrus.redis.enabled"
	RedisHost    = "log.logrus.redis.host"
	RedisKey     = "log.logrus.redis.key"
	RedisFormat  = "log.logrus.redis.format"
	RedisApp     = "log.logrus.redis.app"
	RedisPort    = "log.logrus.redis.port"
	RedisDb      = "log.logrus.redis.db"
	Formatter    = "log.logrus.formatter"
)

func init() {

	log.Println("getting configurations for logrus")

	config.Add(RedisEnabled, false, "enable/disable redis logging")
	config.Add(RedisHost, "localhost", "redis host")
	config.Add(RedisKey, "", "redis key")
	config.Add(RedisFormat, "b2w", "redis format")
	config.Add(RedisApp, "logger", "redis app")
	config.Add(RedisPort, 6379, "redis port")
	config.Add(RedisDb, 0, "redis db")
	config.Add(Formatter, "TEXT", "formatter TEXT/JSON/AWS_CLOUD_WATCH")
}
