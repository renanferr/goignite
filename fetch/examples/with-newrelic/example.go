package main

import (
	"context"
	"fmt"
	"time"

	"github.com/b2wdigital/goignite/v2/fetch"
	"github.com/b2wdigital/goignite/v2/fetch/ext/nrfetch"
)

func main() {
	client := fetch.New()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myOldCtx", "myOldValue")

	o := fetch.Options{
		Url:     "https://pokeapi.co/api/v2/pokemon/ditto",
		Method:  "GET",
		Ctx:     ctx,
		Timeout: time.Duration(1) * time.Second,
	}

	nrfetch.Integrate(&client)

	r := client.Fetch(o)

	if r.Error != nil {
		fmt.Println(r.Error)
	}

	fmt.Println("Result", r.StatusCode, string(r.Body))
}
