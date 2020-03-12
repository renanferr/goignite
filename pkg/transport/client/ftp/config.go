package ftp

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Addr     = "transport.client.ftp.addr"
	Username = "transport.client.ftp.username"
	Password = "transport.client.ftp.password"
	Timeout  = "transport.client.ftp.timeout"
	Retry    = "transport.client.ftp.retry"
)

func init() {

	log.Println("getting configurations for ftp")

	config.Add(Addr, "", "ftp address")
	config.Add(Username, "", "ftp username")
	config.Add(Password, "", "ftp password")
	config.Add(Timeout, 10, "ftp timeout")
	config.Add(Retry, 3, "ftp retry")
}
