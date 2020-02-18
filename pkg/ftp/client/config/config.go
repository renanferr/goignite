package ftp

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const (
	Addr     = "ftp.addr"
	Username = "ftp.username"
	Password = "ftp.password"
	Timeout  = "ftp.timeout"
	Retry    = "ftp.retry"
)

func init() {

	log.Println("getting configurations for ftp")

	config.Add(Addr, "", "ftp address")
	config.Add(Username, "", "ftp username")
	config.Add(Password, "", "ftp password")
	config.Add(Timeout, 10, "ftp timeout")
	config.Add(Retry, 3, "ftp retry")
}
