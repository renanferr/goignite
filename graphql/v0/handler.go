package gigraphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewDefaultHandler(schema *graphql.Schema) *handler.Handler {
	config, _ := DefaultHandlerConfig()
	config.Schema = schema
	h := handler.New(config)

	return h
}
