package tid

import (
	"context"

	"github.com/go-resty/resty/v2"
)

func Middleware(ctx context.Context, client *resty.Client) error {

	if !isEnabled() {
		return nil
	}

	client.OnBeforeRequest(tid)

	return nil
}

func tid(client *resty.Client, request *resty.Request) error {

	ctx := request.Context()

	tidValue := ctx.Value("x-tid").(string)

	request.SetHeader("X-TID", tidValue)

	return nil
}
