package resty

import (
	"context"
	"errors"
	"strconv"
	"strings"

	rootresty "github.com/b2wdigital/goignite/pkg/transport/client/http/resty"
	"github.com/go-resty/resty/v2"
)

type RestyChecker struct {
	client  *resty.Client
	options *rootresty.Options
}

func (c *RestyChecker) Check(ctx context.Context) (err error) {

	request := c.client.R().EnableTrace()

	var response *resty.Response

	response, err = request.Get(strings.Join([]string{c.options.Host, c.options.Health.Endpoint}, ""))

	if response.IsError() {
		return errors.New(strconv.Itoa(response.StatusCode()))
	}

	return err
}

func NewRestyChecker(client *resty.Client, options *rootresty.Options) *RestyChecker {
	return &RestyChecker{client: client, options: options}
}
