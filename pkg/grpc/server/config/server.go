package config

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Port                 = "grpc.server.port"
	MaxConcurrentStreams = "grpc.server.maxconcurrentstreams"
)

func init() {

	log.Println("getting configurations for grpc server")

	config.Add(Port, 9090, "server grpc port")
	config.Add(MaxConcurrentStreams, 5000, "server grpc max concurrent streams")
}
