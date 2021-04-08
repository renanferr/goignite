package cloudevents

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
)

type Server struct {
	handler Handler
	client  client.Client
}

func NewDefault(ctx context.Context, handler Handler) *Server {
	logger := log.FromContext(ctx)
	c, err := v2.NewDefaultClient()
	if err != nil {
		logger.Panic(err.Error())
	}
	return &Server{handler: handler, client: c}
}

func (s *Server) Serve(ctx context.Context) {
	logger := log.FromContext(ctx).WithTypeOf(*s)
	if err := s.client.StartReceiver(ctx, s.handler); err != nil {
		logger.Panic(err.Error())
	}
}
