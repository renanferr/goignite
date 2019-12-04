package main

import (
	"fmt"

	"github.com/jpfaria/goignite/pkg/config"

)

func main() {
	config.Parse()

	fmt.Println(config.Instance.Bool("debug"))
}
