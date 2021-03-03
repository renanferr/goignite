package receiver

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	Url      = "gi.cloudevents.nats.receiver.url"
	Subjects = "gi.cloudevents.nats.receiver.subjects"
)

func init() {

	giconfig.Add(Url, "http://127.0.0.1:4222", "define nats server")
	giconfig.Add(Subjects, []string{"test"}, "define subject")
}

func GetUrl() string {
	return giconfig.String(Url)
}

func GetSubjects() []string {
	return giconfig.Strings(Subjects)
}
