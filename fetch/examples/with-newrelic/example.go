package main

import (
	"context"
	"fmt"
	"time"

	gifetch "github.com/b2wdigital/goignite/v2/fetch"
	gifetchnewrelic "github.com/b2wdigital/goignite/v2/fetch/ext/newrelic/v3"
)

func main() {

	client := gifetch.New()
	client.Use(gifetchnewrelic.New())

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
