package handler

import (
	"net/http"

	"github.com/b2wdigital/goignite/examples/server/http/router/echo/aop/model/response"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/server/http/router/echo"
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

	return echo.JSONResponse(c, http.StatusOK, resp, err)
}
