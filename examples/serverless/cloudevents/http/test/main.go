package main

import (
	"context"
	"log"
	"net/http"

	"github.com/cloudevents/sdk-go"
	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/log/logrus"
	c "github.com/jpfaria/goignite/pkg/serverless/cloudevents/transport/http"
)

type Example struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

func gotEvent(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) error {
	
	log := logrus.FromContext(ctx)
	
	log.Printf("Got Event Context: %+v\n", event.Context)
	data := &Example{}
	if err := event.DataAs(data); err != nil {
		log.Printf("Got Data Error: %s\n", err.Error())
	}
	log.Printf("Got Data: %+v\n", data)
	log.Printf("Got Transport Context: %+v\n", cloudevents.HTTPTransportContextFrom(ctx))
	log.Printf("----------------------------\n")

	if data.Sequence%3 == 0 {
		r := cloudevents.Event{
			Context: cloudevents.EventContextV02{
				Source: *cloudevents.ParseURLRef("/mod3"),
				Type:   "samples.http.mod3",
			}.AsV02(),
			Data: Example{
				Sequence: data.Sequence,
				Message:  "mod 3!",
			},
		}
		resp.RespondWith(200, &r)
		resp.Context = &cloudevents.HTTPTransportResponseContext{
			Header: func() http.Header {
				h := http.Header{}
				h.Set("sample", "magic header")
				h.Set("mod", "3")
				return h
			}(),
		}
		return nil
	}

	return nil
}

func main() {

	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	logrus.Start()

	c.Start(ctx, gotEvent)
}
