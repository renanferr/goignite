package main

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	giecho "github.com/b2wdigital/goignite/echo/v4"
	e "github.com/labstack/echo/v4"
)

func Get(c e.Context) (err error) {

	l := gilog.FromContext(c.Request().Context())

	resp := Response{
		Message: "Hello Google!!",
	}

	err = giconfig.Unmarshal(&resp)
	if err != nil {
		l.Errorf(err.Error())
	}

	return giecho.JSON(c, http.StatusOK, resp, err)
}
