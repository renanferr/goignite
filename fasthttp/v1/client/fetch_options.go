package gifasthttp

import (
	"context"
	"time"

	"github.com/lann/builder"
	"github.com/valyala/fasthttp"
)

type FetchOptions struct {
	Url     string
	Method  string
	Header  *fasthttp.RequestHeader
	Timeout time.Duration
	Body    []byte
	Ctx     context.Context
}

type fetchOptionsBuilder builder.Builder

func (b fetchOptionsBuilder) Url(value string) fetchOptionsBuilder {
	return builder.Set(b, "Url", value).(fetchOptionsBuilder)
}

func (b fetchOptionsBuilder) Method(value string) fetchOptionsBuilder {
	return builder.Set(b, "Method", value).(fetchOptionsBuilder)
}

func (b fetchOptionsBuilder) Header(value *fasthttp.RequestHeader) fetchOptionsBuilder {
	return builder.Set(b, "Header", value).(fetchOptionsBuilder)
}

func (b fetchOptionsBuilder) Timeout(value time.Duration) fetchOptionsBuilder {
	return builder.Set(b, "Timeout", value).(fetchOptionsBuilder)
}

func (b fetchOptionsBuilder) Body(value []byte) fetchOptionsBuilder {
	return builder.Set(b, "Body", value).(fetchOptionsBuilder)
}

func (b fetchOptionsBuilder) Ctx(value context.Context) fetchOptionsBuilder {
	return builder.Set(b, "Ctx", value).(fetchOptionsBuilder)
}

func (b fetchOptionsBuilder) Build() FetchOptions {
	return builder.GetStruct(b).(FetchOptions)
}

var FetchOptionsBuilder = builder.Register(fetchOptionsBuilder{}, FetchOptions{}).(fetchOptionsBuilder)
