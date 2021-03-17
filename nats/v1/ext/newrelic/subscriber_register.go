package newrelic

import (
	"github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
)

func SubscriberRegister(msgHandler nats.MsgHandler) nats.MsgHandler {
	if !IsEnabled() || !newrelic.IsEnabled() {
		return msgHandler
	}

	return nrnats.SubWrapper(newrelic.Application(), msgHandler)
}
