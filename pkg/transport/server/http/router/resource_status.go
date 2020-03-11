package router

import (
	"context"

	"github.com/b2wdigital/goignite/pkg/info"
)

func ResourceStatus(ctx context.Context) ResourceStatusResponse {

	return ResourceStatusResponseBuilder.
		ApplicationName(info.AppName).
		ImplementationBuild(info.BuildVersion).
		ImplementationVersion(info.Version).
		BuildDate(info.BuildDate).
		CommitSHA(info.CommitSHA).
		Build()

}
