package handler

import (
	"context"

	"github.com/cloudevents/sdk-go"
	"github.com/jpfaria/goignite/pkg/log/logrus"
	"github.com/jpfaria/goignite/pkg/serverless/cloudevents/example/model/event"
)

type Response struct {
	Message string
}

func Test2(ctx context.Context, e cloudevents.Event, resp *cloudevents.EventResponse) error {

	log := logrus.FromContext(ctx)

	user := &event.User{}
	if err := e.DataAs(user); err != nil {
		log.Printf("Got Data Error: %s\n", err.Error())
	}

	log.Info(user.Name)

	r := cloudevents.Event{
		Context: cloudevents.EventContextV03{
			Source: *cloudevents.ParseURLRef("/mod3"),
			Type:   "samples.http.mod3",
		}.AsV03(),
		Data: Response{
			Message: "Test 2!!",
		},
	}

	resp.Event = &r

	return nil
}

func Test1(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) error {
	r := cloudevents.Event{
		Context: cloudevents.EventContextV03{
			Source: *cloudevents.ParseURLRef("/mod3"),
			Type:   "samples.http.mod3",
		}.AsV03(),
		Data: Response{
			Message: "Test 1!!",
		},
	}

	resp.Event = &r

	return nil
}
