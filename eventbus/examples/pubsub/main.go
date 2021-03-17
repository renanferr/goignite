package main

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/eventbus"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {
	config.Load()
	logrus.NewLogger()

	eventbus.Start()

	eventbus.Subscribe("exampleint", ExampleInt)
	defer eventbus.Unsubscribe("exampleint", ExampleInt)
	eventbus.Subscribe("examplestring", ExampleString)
	defer eventbus.Unsubscribe("examplestring", ExampleString)

	eventbus.Subscribe("examplestring", ExampleString1)
	defer eventbus.Unsubscribe("examplestring", ExampleString1)
	eventbus.Subscribe("examplestring", ExampleString2)
	defer eventbus.Unsubscribe("examplestring", ExampleString2)

	eventbus.Publish("exampleint", 1)
	eventbus.Publish("examplestring", "Hello World!!")
}

func ExampleInt(i int) {
	log.Infof("logging int %v", i)
}

func ExampleString(m string) {
	log.Infof("logging string %s", m)
}

func ExampleString1(m string) {
	log.Infof("logging string 2 %s", m)
}

func ExampleString2(m string) {
	log.Infof("logging string 1 %s", m)
}
