package girestytid

import (
	"context"

	giinfo "github.com/b2wdigital/goignite/v2/info"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("enabling tid middleware in resty")

	client.OnBeforeRequest(tid)

	logger.Debug("tid middleware successfully enabled in resty")

	return nil
}

func tid(client *resty.Client, request *resty.Request) error {

	ctx := request.Context()

	tidValue, ok := ctx.Value("x-tid").(string)
	if !ok {
		tidValue = giinfo.AppName + "-" + uuid.NewV4().String()
	}

	request.SetHeader("X-TID", tidValue)

	return nil
}
