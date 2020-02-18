package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const Port = "server.grpc.port"
const MaxConcurrentStreams = "server.grpc.maxConcurrentStreams"

func init() {

	log.Println("getting configurations for grpc server")

	config.SetInt(Port, 9090, "server grpc port")
	config.SetUint64(MaxConcurrentStreams, 5000, "server grpc max concurrent streams")
}
