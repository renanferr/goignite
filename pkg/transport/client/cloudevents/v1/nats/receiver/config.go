package receiver

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Url     = "transport.client.cloudevents.nats.receiver.url"
	Subjects = "transport.client.cloudevents.nats.receiver.subjects"
)

func init() {

	log.Println("getting configurations for http server")

	config.Add(Url, "http://127.0.0.1:4222", "define nats server")
	config.Add(Subjects, []string{"test"}, "define subject")
}

func GetUrl() string {
	return config.String(Url)
}

func GetSubjects() []string {
	return config.Strings(Subjects)
}