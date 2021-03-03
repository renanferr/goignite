package gigocloud

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root     = "gi.gocloud"
	resource = root + ".resource"
	tp       = root + ".type"
	region   = root + ".region"
)

func init() {
	giconfig.Add(tp, "memory", "define queue type")
	giconfig.Add(resource, "topicA", "define queue resource")
	giconfig.Add(region, "", "define queue region")
}
