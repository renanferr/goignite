package main

import (
    "context"
    "net/http"

    e "github.com/labstack/echo"

    "github.com/jpfaria/goignite/pkg/config"
    "github.com/jpfaria/goignite/pkg/logging/logrus"
)

func NewHandler() *Handler {
    return &Handler{}
}

type Handler struct {
}

func (u *Handler) Get(c e.Context) error {

    log := logrus.FromContext(context.Background())

    resp := Response{}

    err := config.Unmarshal(&resp)
    if err != nil {
        log.Error(err)
    }

    return c.JSON(http.StatusOK, resp)
}
