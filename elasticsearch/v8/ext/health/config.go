package health

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/elasticsearch/v8"
)

const (
	root        = elasticsearch.ExtRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	config.Add(name, "elasticsearch", "health name")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
