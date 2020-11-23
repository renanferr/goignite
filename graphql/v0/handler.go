package gigraphql

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

const (
	TopicHandler = "topic:gigraphql:handler"
)

func NewDefaultHandler(schema *graphql.Schema) *handler.Handler {
	config, _ := DefaultHandlerConfig()
	config.Schema = schema
	h := handler.New(config)

	gieventbus.Publish(TopicHandler, h)

	return h
}
