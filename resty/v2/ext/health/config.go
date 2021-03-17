package health

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/resty/v2"
)

const (
	root        = resty.ExtRoot + ".health"
	name        = root + ".name"
	host        = root + ".host"
	endpoint    = root + ".endpoint"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {

	config.Add(name, "rest api", "health name")
	config.Add(host, "", "health host")
	config.Add(endpoint, "/resource-status", "health host")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
