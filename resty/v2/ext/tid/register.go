package tid

import (
	"context"

	"github.com/b2wdigital/goignite/v2/info"
	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	client.OnBeforeRequest(tid)

	return nil
}

func tid(client *resty.Client, request *resty.Request) error {

	ctx := request.Context()

	tidValue, ok := ctx.Value("x-tid").(string)
	if !ok {
		tidValue = info.AppName + "-" + uuid.NewV4().String()
	}

	request.SetHeader("X-TID", tidValue)

	return nil
}
