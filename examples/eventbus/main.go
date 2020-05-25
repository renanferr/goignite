package main

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
)

func main() {
	giconfig.Load()
	gilogrus.NewLogger()

	gieventbus.Subscribe("exampleint", ExampleInt)
	gieventbus.Subscribe("examplestring", ExampleString)

	gieventbus.Publish("exampleint", 1)
	gieventbus.Publish("examplestring", "Hello World!!")
}

func ExampleInt(i int) {
	gilog.Infof("logging int %v", i)
}

func ExampleString(m string) {
	gilog.Infof("logging string %s", m)
}