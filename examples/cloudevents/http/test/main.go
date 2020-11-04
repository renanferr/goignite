package main

import (
	"context"
	"net/http"

	chttp "github.com/b2wdigital/goignite/cloudevents/v1/http"
	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
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

	giconfig.Load()

	ctx := context.Background()

	gilogrus.NewLogger()

	chttp.Start(ctx, Test2, "POST")
}

func Test2(ctx context.Context, e cloudevents.Event, resp *cloudevents.EventResponse) error {

	l := gilog.FromContext(ctx)

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
