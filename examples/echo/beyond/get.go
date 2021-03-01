package main

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
	gilog "github.com/b2wdigital/goignite/log"
	e "github.com/labstack/echo/v4"
)

func Get(c e.Context) (err error) {

	logger := gilog.FromContext(c.Request().Context())

	resp := Response{
		Message: "Hello Google!!",
	}

	err = giconfig.Unmarshal(&resp)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return giecho.JSON(c, http.StatusOK, resp, err)
}
