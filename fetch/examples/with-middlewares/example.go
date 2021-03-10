package main

import (
	"context"
	"fmt"
	"time"

	"github.com/b2wdigital/goignite/v2/fetch"
)

func main() {
	client := fetch.New()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myOldCtx", "myOldValue")

	o := fetch.Options{
		Url:     "http://product-v3-americanas-npf.internal.b2w.io/product/1264011424",
		Method:  "GET",
		Ctx:     ctx,
		Timeout: time.Duration(1) * time.Second,
	}

	client.OnBeforeRequest(func(o fetch.Options, ctx context.Context) context.Context {
		fmt.Println("ctx before 1 - 1", ctx.Value("myOldCtx"))
		ctx = context.WithValue(ctx, "myNewCtx", "myNewValue")
		return ctx
	})
	client.OnBeforeRequest(func(o fetch.Options, ctx context.Context) context.Context {
		fmt.Println("ctx before 2 - 1", ctx.Value("myOldCtx"))
		ctx = context.WithValue(ctx, "myNewCtx", "myNewValue")
		return ctx
	})

	client.OnAfterRequest(func(o fetch.Options, r fetch.Response, ctx context.Context) {
		fmt.Println("ctx after 1", ctx.Value("myOldCtx"))
		fmt.Println("ctx after 2", ctx.Value("myNewCtx"))
	})

	r := client.Fetch(o)

	if r.Error != nil {
		fmt.Println(r.Error)
	}

	fmt.Println("Result", r.StatusCode, string(r.Body))
}
