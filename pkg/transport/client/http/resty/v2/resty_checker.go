package resty

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

type RestyChecker struct {
	client  *resty.Client
	options *Options
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

func NewRestyChecker(client *resty.Client, options *Options) *RestyChecker {
	return &RestyChecker{client: client, options: options}
}
