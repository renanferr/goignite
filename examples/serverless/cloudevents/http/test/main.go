package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cloudevents/sdk-go"
	"github.com/go-playground/validator/v10"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/b2wdigital/goignite/pkg/serverless/cloudevents/example/handler"
	"github.com/b2wdigital/goignite/pkg/serverless/cloudevents/example/model/event"
	c "github.com/b2wdigital/goignite/pkg/serverless/cloudevents/transport/http"
)

type Example struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

func Test2(ctx context.Context, e cloudevents.Event, resp *cloudevents.EventResponse) error {

	l := logrus.FromContext(ctx)

	user := &event.User{}
	if err := e.DataAs(user); err != nil {
		l.Printf("Got Data Error: %s\n", err.Error())
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {

		resp.Status = http.StatusUnprocessableEntity
		resp.Event = &cloudevents.Event{
			Context: cloudevents.EventContextV03{
				Source: *cloudevents.ParseURLRef("/mod3"),
				Type:   "samples.http.mod3",
			}.AsV03(),
		}

		return err
	}

	resp.Status = http.StatusCreated

	return handler.Test2(ctx, e, resp)
}

func main() {

	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	logrus.Start()

	c.Start(ctx, Test2, "POST")
}
