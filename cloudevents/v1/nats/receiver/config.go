package receiver

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Url      = "gi.cloudevents.nats.receiver.url"
	Subjects = "gi.cloudevents.nats.receiver.subjects"
)

func init() {

	log.Println("getting configurations for http server")

	giconfig.Add(Url, "http://127.0.0.1:4222", "define nats server")
	giconfig.Add(Subjects, []string{"test"}, "define subject")
}

func GetUrl() string {
	return giconfig.String(Url)
}

func GetSubjects() []string {
	return giconfig.Strings(Subjects)
}
