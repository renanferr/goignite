package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	gifasthttp "github.com/b2wdigital/goignite/v2/fasthttp/v1/client"
)

func main() {

	ctx := context.Background()

	client := gifasthttp.NewDefaultFetch(ctx)
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
