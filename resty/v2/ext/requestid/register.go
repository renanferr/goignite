package girestyrequestid

import (
	"context"

	"github.com/go-resty/resty/v2"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	client.OnBeforeRequest(requestId)

	return nil
}

func requestId(client *resty.Client, request *resty.Request) error {

	ctx := request.Context()

	idValue, ok := ctx.Value("requestId").(string)
	if !ok {
		idValue = uuid.NewV4().String()
	}

	request.SetHeader("X-Request-ID", idValue)

	return nil
}
