package main

import (
	"context"
	"fmt"
	"time"

	"github.com/b2wdigital/goignite/v2/fetch"
)

func main() {
	client := fetch.New()

	o := fetch.Options{
		Url:     "https://pokeapi.co/api/v2/pokemon/ditto",
		Method:  "GET",
		Ctx:     context.Background(),
		Timeout: time.Duration(1) * time.Second,
	}

	r := client.Fetch(o)

	if r.Error != nil {
		fmt.Println(r.Error)
	}

	fmt.Println("Result", r.StatusCode, string(r.Body))
}
