package gilogrus

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root         = "gi.log.logrus"
	redisRoot    = root + ".redis"
	redisEnabled = redisRoot + ".enabled"
	redisHost    = redisRoot + ".host"
	redisKey     = redisRoot + ".key"
	redisFormat  = redisRoot + ".format"
	redisApp     = redisRoot + ".app"
	redisPort    = redisRoot + ".port"
	redisDb      = redisRoot + ".db"
	formatter    = root + ".formatter"
)

func init() {

	giconfig.Add(redisEnabled, false, "enable/disable redis logging")
	giconfig.Add(redisHost, "localhost", "redis host")
	giconfig.Add(redisKey, "", "redis key")
	giconfig.Add(redisFormat, "v1", "redis format")
	giconfig.Add(redisApp, "logger", "redis app")
	giconfig.Add(redisPort, 6379, "redis port")
	giconfig.Add(redisDb, 0, "redis db")
	giconfig.Add(formatter, "TEXT", "formatter TEXT/JSON/AWS_CLOUD_WATCH")
}
