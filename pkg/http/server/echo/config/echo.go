package config

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const HideBanner = "http.server.echo.hidebanner"

func init() {

	log.Println("getting configurations for echo")

	config.Add(HideBanner, true, "echo hide/show banner")

}

func GetHideBanner() bool {
	return config.Bool(HideBanner)
}