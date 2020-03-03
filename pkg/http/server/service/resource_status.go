package service

import (
	"context"

	"github.com/jpfaria/goignite/pkg/http/server/model/response"
	"github.com/jpfaria/goignite/pkg/info"
)

func ResourceStatus(ctx context.Context) response.ResourceStatusResponse {

	return response.ResourceStatusResponseBuilder.
		ApplicationName(info.AppName).
		ImplementationBuild(info.BuildVersion).
		ImplementationVersion(info.Version).
		BuildDate(info.BuildDate).
		CommitSHA(info.CommitSHA).
		Build()

}
