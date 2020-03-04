package service

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/http/router/model/response"
	"github.com/b2wdigital/goignite/pkg/info"
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
