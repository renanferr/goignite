package nats

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Server  = "transport.client.cloudevents.nats.server"
	Subject = "transport.client.cloudevents.nats.subject"
)

func init() {

	log.Println("getting configurations for http server")

	config.Add(Server, "nats://127.0.0.1:4222", "define nats server")
	config.Add(Subject, "test", "define subject")
}

func GetServer() string {
	return config.String(Server)
}

func GetSubject() string {
	return config.String(Subject)
}