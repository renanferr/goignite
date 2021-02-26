package gilogrus

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	RedisEnabled = "gi.log.logrus.redis.enabled"
	RedisHost    = "gi.log.logrus.redis.host"
	RedisKey     = "gi.log.logrus.redis.key"
	RedisFormat  = "gi.log.logrus.redis.format"
	RedisApp     = "gi.log.logrus.redis.app"
	RedisPort    = "gi.log.logrus.redis.port"
	RedisDb      = "gi.log.logrus.redis.db"
	Formatter    = "gi.log.logrus.formatter"
)

func init() {

	giconfig.Add(RedisEnabled, false, "enable/disable redis logging")
	giconfig.Add(RedisHost, "localhost", "redis host")
	giconfig.Add(RedisKey, "", "redis key")
	giconfig.Add(RedisFormat, "v1", "redis format")
	giconfig.Add(RedisApp, "logger", "redis app")
	giconfig.Add(RedisPort, 6379, "redis port")
	giconfig.Add(RedisDb, 0, "redis db")
	giconfig.Add(Formatter, "TEXT", "formatter TEXT/JSON/AWS_CLOUD_WATCH")
}
