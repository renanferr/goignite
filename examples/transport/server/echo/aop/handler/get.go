package handler

import (
	"net/http"

	"github.com/b2wdigital/goignite/examples/transport/server/echo/aop/model/response"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/server/echo/v4"
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

	return echo.JSON(c, http.StatusOK, resp, err)
}
