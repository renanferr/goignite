package ftp

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const (
	Addr     = "ftp.client.addr"
	Username = "ftp.client.username"
	Password = "ftp.client.password"
	Timeout  = "ftp.client.timeout"
	Retry    = "ftp.client.retry"
)

func init() {

	log.Println("getting configurations for ftp")

	config.Add(Addr, "", "ftp address")
	config.Add(Username, "", "ftp username")
	config.Add(Password, "", "ftp password")
	config.Add(Timeout, 10, "ftp timeout")
	config.Add(Retry, 3, "ftp retry")
}
