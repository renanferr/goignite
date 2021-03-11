package main

import (
	"context"
	"fmt"
	"time"

	gifasthttp "github.com/b2wdigital/goignite/v2/fasthttp/v1/client"
	gifetch "github.com/b2wdigital/goignite/v2/fetch"
)

func main() {

	ctx := context.Background()

	fasthttpClient := gifasthttp.NewDefaultClient(ctx)

	client := gifetch.New(fasthttpClient)
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
