package gigraphql

import (
	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/graphql-go/handler"
)

const (
	rootConfig = "gi.graphql"

	handlerConfig = rootConfig + ".handler"

	pretty           = handlerConfig + ".pretty"
	enableGraphiQL   = handlerConfig + ".graphiQL"
	enablePlayground = handlerConfig + ".playground"
)

func init() {
	giconfig.Add(pretty, false, "enable/disable pretty print")
	giconfig.Add(enableGraphiQL, false, "enable/disable GraphiQL")
	giconfig.Add(enablePlayground, true, "enable/disable Playground")
}

func DefaultHandlerConfig() (*handler.Config, error) {

	o := &handler.Config{}

	err := giconfig.UnmarshalWithPath(handlerConfig, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
