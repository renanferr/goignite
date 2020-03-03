package handler

import (
	"context"

	"github.com/cloudevents/sdk-go"
	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/log/logrus"
	"github.com/jpfaria/goignite/pkg/serverless/cloudevents/example/model/event"
)

var (
	options *Response
)

func init() {
	config.Add("local.message", "gen", "generator output path")
}


type Response struct {
	Message string
}

func Start(ctx context.Context) {
	log := logrus.FromContext(ctx)

	log.Info("starting application")

	options = new(Response)

	err := config.UnmarshalWithPath("local", &options)
	if err != nil {
		log.Error(err)
	}
}

func Stop(ctx context.Context) {
	log := logrus.FromContext(ctx)
	log.Info("stopping application")
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
			Message: "Test 3!!",
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
