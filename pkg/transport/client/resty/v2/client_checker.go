package resty

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

type ClientChecker struct {
	client  *resty.Client
	options *Options
}

func (c *ClientChecker) Check(ctx context.Context) (err error) {

	request := c.client.R().EnableTrace()

	var response *resty.Response

	response, err = request.Get(strings.Join([]string{c.options.Host, c.options.Health.Endpoint}, ""))

	if response.IsError() {
		return errors.New(strconv.Itoa(response.StatusCode()))
	}

	return err
}

func NewClientChecker(client *resty.Client, options *Options) *ClientChecker {
	return &ClientChecker{client: client, options: options}
}
