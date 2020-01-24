package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const HideBanner = "server.http.echo.hidebanner"

func init() {

	log.Println("getting configurations for echo")

	config.Add(HideBanner, true, "echo hide/show banner")

}


func GetHideBanner() bool {
	return config.Instance.Bool(HideBanner)
}