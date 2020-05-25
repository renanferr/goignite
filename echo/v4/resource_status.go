package giecho

import (
	"net/http"

	"github.com/b2wdigital/goignite/pkg/rest/response"
	"github.com/labstack/echo/v4"
)

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get(c echo.Context) error {
	return JSON(c, http.StatusOK, response.NewResourceStatus(), nil)
}
