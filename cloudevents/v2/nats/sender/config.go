package sender

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	Url = "gi.cloudevents.nats.sender.url"
)

func init() {

	giconfig.Add(Url, "http://127.0.0.1:4222", "define nats server")
}

func GetUrl() string {
	return giconfig.String(Url)
}
