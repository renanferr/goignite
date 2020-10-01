package gieventbus

import (
	evbus "github.com/asaskevich/EventBus"
)

var (
	bus evbus.Bus
)

func init() {
	bus = evbus.New()
}

func Subscribe(topic string, fn interface{}) error {
	return bus.Subscribe(topic, fn)
}

func SubscribeOnce(topic string, fn interface{}) error {
	return bus.SubscribeOnce(topic, fn)
}

func SubscribeOnceAsync(topic string, fn interface{}) error {
	return bus.SubscribeOnceAsync(topic, fn)
}

func Publish(topic string, args ...interface{}) {
	bus.Publish(topic, args...)
}

func Unsubscribe(topic string, fn interface{}) error {
	return bus.Unsubscribe(topic, fn)
}

func WaitAsync() {
	bus.WaitAsync()
}
