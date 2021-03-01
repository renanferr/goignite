package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gigocloud "github.com/b2wdigital/goignite/gocloud/v0"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	"gocloud.dev/pubsub"
)

func main() {

	giconfig.Load()

	ctx := context.Background()

	gilogrus.NewLogger()

	logger := gilog.FromContext(ctx)

	topic, err := gigocloud.NewDefaultTopic(ctx)
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
