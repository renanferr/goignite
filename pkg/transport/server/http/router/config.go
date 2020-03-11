package router

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Port        = "http.server.port"
	StatusRoute = "http.server.route.status"
	HealthRoute = "http.server.route.health"
)

func init() {

	log.Println("getting configurations for http server")

	config.Add(Port, 8080, "server http port")
	config.Add(StatusRoute, "/resource-status", "define status url")
	config.Add(HealthRoute, "/health", "define health url")

}

func GetPort() int {
	return config.Int(Port)
}

func GetStatusRoute() string {
	return config.String(StatusRoute)
}

func GetHealthRoute() string {
	return config.String(HealthRoute)
}
