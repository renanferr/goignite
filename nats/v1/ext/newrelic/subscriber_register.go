package ginatsnewrelic

import (
	ginewrelic "github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
)

func SubscriberRegister(msgHandler nats.MsgHandler) nats.MsgHandler {
	if !IsEnabled() || !ginewrelic.IsEnabled() {
		return msgHandler
	}

	return nrnats.SubWrapper(ginewrelic.Application(), msgHandler)
}
