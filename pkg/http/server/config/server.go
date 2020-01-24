package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const Port = "server.http.port"
const StatusRoute = "server.http.route.status"
const HealthRoute = "server.http.route.health"

func init() {

	log.Println("getting configurations for http server")

	config.Add(Port, 8080, "server http port")
	config.Add(StatusRoute, "/resource-status", "define status url")
	config.Add(HealthRoute, "/health", "define health url")

}

func GetPort() int {
	return config.Instance.Int(Port)
}

func GetStatusRoute() string {
	return config.Instance.String(StatusRoute)
}

func GetHealthRoute() string {
	return config.Instance.String(HealthRoute)
}