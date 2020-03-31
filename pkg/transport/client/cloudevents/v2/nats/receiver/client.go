package receiver

import (
	l "github.com/b2wdigital/goignite/pkg/log"
	natsce "github.com/cloudevents/sdk-go/v2/protocol/nats"
	"github.com/nats-io/nats.go"
)

func NewSender(conn *nats.Conn, subject string) (*natsce.Sender, error) {

	sender, err := natsce.NewSenderFromConn(conn, subject)
	if err != nil {
		l.Fatalf("failed to create transport: %s", err.Error())
	}

	return sender, nil
}
