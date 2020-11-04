package sender

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Url = "gi.cloudevents.nats.sender.url"
)

func init() {

	log.Println("getting configurations for http server")

	giconfig.Add(Url, "http://127.0.0.1:4222", "define nats server")
}

func GetUrl() string {
	return giconfig.String(Url)
}
