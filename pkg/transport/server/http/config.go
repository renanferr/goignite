package http

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Port        = "transport.server.http.port"
	StatusRoute = "transport.server.http.route.status"
	HealthRoute = "transport.server.http.route.health"
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
