package handler

import (
	"net/http"

	"github.com/jpfaria/goignite/pkg/http/server/model/response"
	"github.com/jpfaria/goignite/pkg/info"
	"github.com/labstack/echo"
)

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get(c echo.Context) error {

	resourceStatus := response.ResourceStatusResponseBuilder.
		ApplicationName(info.AppName).
		ImplementationBuild(info.BuildVersion).
		ImplementationVersion(info.Version).
		BuildDate(info.BuildDate).
		CommitSHA(info.CommitSHA).
		Build()

	return c.JSON(http.StatusOK, resourceStatus)
}
