package sender

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	cloudevents "github.com/cloudevents/sdk-go"
)

type Sender struct {
	options *Options
	subject string
}

func NewSender(subject string, options *Options) *Sender {
	return &Sender{subject: subject, options: options}
}

func (s *Sender) Send(ctx context.Context, event cloudevents.Event) error {

	l := gilog.FromContext(ctx)

	c, err := NewClient(s.options.Url, s.subject)
	if err != nil {
		l.Errorf("failed to create client: %s", err.Error())
		return err
	}

	l.Debugf("sending message to the %s on subject %s", s.options.Url, s.subject)

	_, ev, err := c.Send(ctx, event)
	if err != nil {
		l.Errorf("failed to sent event %s. %s", ev.Type(), err.Error())
		return err
	}

	l.Infof("message sent to the %s for subject %s", s.options.Url, s.subject)

	return nil

}

func NewDefaultSender(subject string) *Sender {
	l := gilog.FromContext(context.Background())

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewSender(subject, o)
}
