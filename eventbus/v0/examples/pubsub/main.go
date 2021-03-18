package main

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {
	config.Load()
	logrus.NewLogger()

	v0.Start()

	v0.Subscribe("exampleint", ExampleInt)
	defer v0.Unsubscribe("exampleint", ExampleInt)
	v0.Subscribe("examplestring", ExampleString)
	defer v0.Unsubscribe("examplestring", ExampleString)

	v0.Subscribe("examplestring", ExampleString1)
	defer v0.Unsubscribe("examplestring", ExampleString1)
	v0.Subscribe("examplestring", ExampleString2)
	defer v0.Unsubscribe("examplestring", ExampleString2)

	v0.Publish("exampleint", 1)
	v0.Publish("examplestring", "Hello World!!")
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
