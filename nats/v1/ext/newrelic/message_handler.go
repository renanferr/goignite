package newrelic

import (
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
)

func MessageHandler(msgHandler nats.MsgHandler) nats.MsgHandler {
	if !isEnabled() {
		return msgHandler
	}

	return nrnats.SubWrapper(ginewrelic.Application(), msgHandler)
}
