package sender

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Url = "transport.client.cloudevents.nats.sender.url"
)

func init() {

	log.Println("getting configurations for http server")

	config.Add(Url, "http://127.0.0.1:4222", "define nats server")
}

func GetUrl() string {
	return config.String(Url)
}
