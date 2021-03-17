package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/gocloud/v0"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
	"gocloud.dev/pubsub"
)

func main() {

	config.Load()

	ctx := context.Background()

	logrus.NewLogger()

	logger := log.FromContext(ctx)

	topic, err := gocloud.NewDefaultTopic(ctx)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	meta := map[string]string{}

	data := []byte("Hello, World!")

	message := &pubsub.Message{
		Body:     data,
		Metadata: meta,
	}

	if err := topic.Send(ctx, message); err != nil {
		logger.Fatalf(err.Error())
	}

	defer topic.Shutdown(ctx)

	logger.Infof("sucesss message send")

	// Don't works using memory
	// subscription, err := gocloud.NewDefaultSubscription(ctx)
	// if err != nil {
	// 	logger.Fatalf(err.Error())
	// }

	// Loop on received messages.
	// for {
	// 	m, err := subscription.Receive(ctx)
	// 	if err != nil {
	// 		logger.Info("Receiving message: %v", err)
	// 		break
	// 	}
	// 	logger.Info("Got message: ", string(m.Body))
	// 	m.Ack()
	// }

	// defer subscription.Shutdown(ctx)
}
