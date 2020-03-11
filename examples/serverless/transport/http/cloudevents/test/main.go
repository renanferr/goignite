package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/log/logrus/v1"
	cloudevents2 "github.com/b2wdigital/goignite/pkg/transport/serverless/transport/http/cloudevents"
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

func main() {

	var err error

	err = config.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	log.NewLogger(v1.NewLogger())

	cloudevents2.Start(ctx, Test2, "POST")
}

func Test2(ctx context.Context, e cloudevents.Event, resp *cloudevents.EventResponse) error {

	l := log.FromContext(ctx)

	user := &User{}
	if err := e.DataAs(user); err != nil {
		l.Errorf("Got Data Error: %s\n", err.Error())
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {

		re := cloudevents.NewEvent()
		re.SetSource("/mod1")
		re.SetType("samples.http.mod3")

		resp.Status = http.StatusUnprocessableEntity
		resp.Event = &re

		return err
	}

	resp.Status = http.StatusCreated

	r := cloudevents.NewEvent()
	r.SetData(Example{Message: "Test 3!!"})
	r.SetSource("/mod1")
	r.SetType("samples.http.mod3")

	resp.Event = &r

	return nil
}
