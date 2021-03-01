package health

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gielasticsearch "github.com/b2wdigital/goignite/elasticsearch/v8"
)

const (
	root        = gielasticsearch.ExtRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	giconfig.Add(name, "elasticsearch", "health name")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
