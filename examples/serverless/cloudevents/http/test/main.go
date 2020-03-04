package main

import (
	"context"
	"log"
	"net/http"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	c "github.com/b2wdigital/goignite/pkg/serverless/cloudevents/transport/http"
	"github.com/cloudevents/sdk-go"
	"github.com/go-playground/validator/v10"
)

type Example struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	CPF   string `json:"cpf"`
}

func Test2(ctx context.Context, e cloudevents.Event, resp *cloudevents.EventResponse) error {

	l := logrus.FromContext(ctx)

	user := &User{}
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

	r := cloudevents.Event{
		Context: cloudevents.EventContextV03{
			Source: *cloudevents.ParseURLRef("/mod3"),
			Type:   "samples.http.mod3",
		}.AsV03(),
		Data: Example{
			Message: "Test 3!!",
		},
	}

	resp.Event = &r

	return nil
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
