package handler

import (
	"net/http"

	"github.com/b2wdigital/goignite/examples/transport/server/http/router/echo/aop/model/response"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/server/http/router/echo/v4"
	e "github.com/labstack/echo/v4"
)

func Get(c e.Context) (err error) {

	l := log.FromContext(c.Request().Context())

	resp := response.Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Errorf(err.Error())
	}

	return v4.JSONResponse(c, http.StatusOK, resp, err)
}
