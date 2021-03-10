package main

import (
	"context"
	"fmt"
	"time"

	gifetch "github.com/b2wdigital/goignite/v2/fetch"
)

type ExampleMiddleware1 struct {
}

func (e ExampleMiddleware1) OnBeforeRequest(ctx context.Context, options gifetch.Options) context.Context {
	fmt.Println("ctx before 1 - 1", ctx.Value("myOldCtx"))
	ctx = context.WithValue(ctx, "myNewCtx", "myNewValue")
	return ctx
}

func (e ExampleMiddleware1) OnAfterRequest(ctx context.Context, options gifetch.Options, response gifetch.Response) {
	fmt.Println("ctx after 1", ctx.Value("myOldCtx"))
	fmt.Println("ctx after 2", ctx.Value("myNewCtx"))
}

func NewExampleMiddleware1() gifetch.Middleware {
	return &ExampleMiddleware1{}
}

func main() {
	client := gifetch.New()
	client.Use(NewExampleMiddleware1())

	client.OnBeforeRequest(func(ctx context.Context, options gifetch.Options) context.Context {
		fmt.Println("ctx before 2 - 1", ctx.Value("myOldCtx"))
		ctx = context.WithValue(ctx, "myNewCtx", "myNewValue")
		return ctx
	})

	ctx := context.Background()
	ctx = context.WithValue(ctx, "myOldCtx", "myOldValue")

	o := gifetch.Options{
		Url:     "http://product-v3-americanas-npf.internal.b2w.io/product/1264011424",
		Method:  "GET",
		Ctx:     ctx,
		Timeout: time.Duration(1) * time.Second,
	}

	r := client.Fetch(o)

	if r.Error != nil {
		fmt.Println(r.Error)
	}

	fmt.Println("Result", r.StatusCode, string(r.Body))
}
