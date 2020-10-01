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
	defer gieventbus.Unsubscribe("exampleint", ExampleInt)
	gieventbus.Subscribe("examplestring", ExampleString)
	defer gieventbus.Unsubscribe("examplestring", ExampleString)

	gieventbus.Subscribe("examplestring", ExampleString1)
	defer gieventbus.Unsubscribe("examplestring", ExampleString1)
	gieventbus.Subscribe("examplestring", ExampleString2)
	defer gieventbus.Unsubscribe("examplestring", ExampleString2)

	gieventbus.Publish("exampleint", 1)
	gieventbus.Publish("examplestring", "Hello World!!")
}

func ExampleInt(i int) {
	gilog.Infof("logging int %v", i)
}

func ExampleString(m string) {
	gilog.Infof("logging string %s", m)
}

func ExampleString1(m string) {
	gilog.Infof("logging string 2 %s", m)
}

func ExampleString2(m string) {
	gilog.Infof("logging string 1 %s", m)
}
