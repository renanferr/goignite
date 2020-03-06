package handler

import (
	"net/http"

	"github.com/b2wdigital/goignite/examples/http/router/echo/aop/model/response"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/http/router/echo/parser"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	e "github.com/labstack/echo/v4"
)

func Get(c e.Context) (err error) {

	l := logrus.FromContext(c.Request().Context())

	resp := response.Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Error(err)
	}

	return parser.JSONResponse(c, http.StatusOK, resp, err)
}
