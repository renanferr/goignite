package kafka

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root      = "gi.kafka"
	address   = root + ".address"
	topic     = root + ".topic"
	partition = root + ".partition"
	network   = root + ".network"
)

func init() {
	config.Add(address, "localhost:9092", "defines host address")
	config.Add(topic, "", "defines topic name")
	config.Add(partition, 0, "defines partition number")
	config.Add(network, "tcp", "defines network protocol")
}
