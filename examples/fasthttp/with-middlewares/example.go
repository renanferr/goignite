package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	gifasthttp "github.com/b2wdigital/goignite/v2/fasthttp/v1/client"
)

type ExampleMiddleware1 struct {
}

func (e ExampleMiddleware1) OnBeforeRequest(ctx context.Context, options gifasthttp.FetchOptions) context.Context {
	fmt.Println("ctx before 1 - 1", ctx.Value("myOldCtx"))
	ctx = context.WithValue(ctx, "myNewCtx", "myNewValue")
	return ctx
}

func (e ExampleMiddleware1) OnAfterRequest(ctx context.Context, options gifasthttp.FetchOptions, response gifasthttp.Response) {
	fmt.Println("ctx after 1", ctx.Value("myOldCtx"))
	fmt.Println("ctx after 2", ctx.Value("myNewCtx"))
}

func NewExampleMiddleware1() gifasthttp.Middleware {
	return &ExampleMiddleware1{}
}

func main() {

	ctx := context.Background()

	client := gifasthttp.NewDefaultFetch(ctx)

	client.Use(NewExampleMiddleware1())

	client.OnBeforeRequest(func(ctx context.Context, options gifasthttp.FetchOptions) context.Context {
		fmt.Println("ctx before 2 - 1", ctx.Value("myOldCtx"))
		ctx = context.WithValue(ctx, "myNewCtx", "myNewValue")
		return ctx
	})

	ctx = context.WithValue(ctx, "myOldCtx", "myOldValue")

	opt := gifasthttp.FetchOptionsBuilder.
		Url("http://product-v3-americanas-npf.internal.b2w.io/product/1264011424").
		Method(http.MethodGet).
		Ctx(ctx).
		Timeout(time.Duration(1) * time.Second).
		Build()

	r := client.Fetch(opt)

	if r.Error != nil {
		fmt.Println(r.Error)
	}

	fmt.Println("Result", r.StatusCode, string(r.Body))
}
