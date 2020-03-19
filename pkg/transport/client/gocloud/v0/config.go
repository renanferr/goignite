package gocloud

import (
	"github.com/b2wdigital/goignite/pkg/config"
)

// configs ..
const (
	Resource = "transport.client.gocloud.resource"
	Type     = "transport.client.gocloud.type"
	Region   = "transport.client.gocloud.region"
)

func init() {
	config.Add(Type, "memory", "define queue type")
	config.Add(Resource, "topicA", "define queue resource")
	config.Add(Region, "", "define queue region")
}
