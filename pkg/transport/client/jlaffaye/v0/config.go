package jlaffaye

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Addr     = "transport.client.jlaffaye.addr"
	Username = "transport.client.jlaffaye.username"
	Password = "transport.client.jlaffaye.password"
	Timeout  = "transport.client.jlaffaye.timeout"
	Retry    = "transport.client.jlaffaye.retry"
)

func init() {

	log.Println("getting configurations for ftp")

	config.Add(Addr, "", "ftp address")
	config.Add(Username, "", "ftp username")
	config.Add(Password, "", "ftp password")
	config.Add(Timeout, 10, "ftp timeout")
	config.Add(Retry, 3, "ftp retry")
}
