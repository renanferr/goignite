package main

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus/v1"
	"github.com/b2wdigital/goignite/pkg/transport/client/gocloud/v0"
	"gocloud.dev/pubsub"
)

func main() {

	config.Load()

	ctx := context.Background()

	log.NewLogger(logrus.NewLogger())

	l := log.FromContext(ctx)

	topic, err := gocloud.NewDefaultTopic(ctx)
	if err != nil {
		l.Fatalf(err.Error())
	}

	meta := map[string]string{}

	data := []byte("Hello, World!")

	message := &pubsub.Message{
		Body:     data,
		Metadata: meta,
	}

	if err := topic.Send(ctx, message); err != nil {
		l.Fatalf(err.Error())
	}

	defer topic.Shutdown(ctx)

	l.Infof("sucesss message send")

	// Don't works using memory
	// subscription, err := gocloud.NewDefaultSubscription(ctx)
	// if err != nil {
	// 	l.Fatalf(err.Error())
	// }

	// Loop on received messages.
	// for {
	// 	m, err := subscription.Receive(ctx)
	// 	if err != nil {
	// 		l.Info("Receiving message: %v", err)
	// 		break
	// 	}
	// 	l.Info("Got message: ", string(m.Body))
	// 	m.Ack()
	// }

	// defer subscription.Shutdown(ctx)
}
