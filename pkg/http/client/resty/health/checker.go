package health

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/jpfaria/goignite/pkg/http/client/resty/model"
)

type RestyChecker struct {
	client  *resty.Client
	options *model.Options
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

func NewRestyChecker(client *resty.Client, options *model.Options) *RestyChecker {
	return &RestyChecker{client: client, options: options}
}
