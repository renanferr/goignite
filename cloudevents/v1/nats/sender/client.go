package sender

import (
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	cloudeventsnats "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/nats"
)

func NewClient(url string, subject string) (client.Client, error) {

	t, err := cloudeventsnats.New(url, subject)
	if err != nil {
		gilog.Fatalf("failed to create transport: %s", err.Error())
	}

	return client.New(t)
}
