package receiver

import (
	gilog "github.com/b2wdigital/goignite/log"
	natsce "github.com/cloudevents/sdk-go/protocol/nats/v2"
	"github.com/nats-io/nats.go"
)

func NewSender(conn *nats.Conn, subject string) (*natsce.Sender, error) {

	sender, err := natsce.NewSenderFromConn(conn, subject)
	if err != nil {
		gilog.Fatalf("failed to create transport: %s", err.Error())
	}

	return sender, nil
}
