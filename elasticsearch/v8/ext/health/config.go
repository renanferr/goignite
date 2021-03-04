package gielasticsearchhealth

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gielasticsearch "github.com/b2wdigital/goignite/v2/elasticsearch/v8"
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
