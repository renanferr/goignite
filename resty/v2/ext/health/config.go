package girestyhealth

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giresty "github.com/b2wdigital/goignite/v2/resty/v2"
)

const (
	root        = giresty.ExtRoot + ".health"
	name        = root + ".name"
	host        = root + ".host"
	endpoint    = root + ".endpoint"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {

	giconfig.Add(name, "rest api", "health name")
	giconfig.Add(host, "", "health host")
	giconfig.Add(endpoint, "/resource-status", "health host")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
