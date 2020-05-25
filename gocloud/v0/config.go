package gigocloud

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

// configs ..
const (
	Resource = "gi.gocloud.resource"
	Type     = "gi.gocloud.type"
	Region   = "gi.gocloud.region"
)

func init() {
	giconfig.Add(Type, "memory", "define queue type")
	giconfig.Add(Resource, "topicA", "define queue resource")
	giconfig.Add(Region, "", "define queue region")
}
