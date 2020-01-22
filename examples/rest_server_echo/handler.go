package main

import (
	"net/http"

	"github.com/jpfaria/goignite/pkg/http/server/echo/parser"
	e "github.com/labstack/echo"

	"github.com/jpfaria/goignite/pkg/config"
)

func NewHandler() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (u *Handler) Get(c e.Context) (err error) {

	resp := Response{}

	err = config.Unmarshal(&resp)

	return parser.JSONResponse(c, http.StatusOK, resp, err)
}
