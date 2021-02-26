package gijlaffaye

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Addr     = "gi.jlaffaye.addr"
	Username = "gi.jlaffaye.username"
	Password = "gi.jlaffaye.password"
	Timeout  = "gi.jlaffaye.timeout"
	Retry    = "gi.jlaffaye.retry"
)

func init() {

	giconfig.Add(Addr, "", "ftp address")
	giconfig.Add(Username, "", "ftp username")
	giconfig.Add(Password, "", "ftp password")
	giconfig.Add(Timeout, 10, "ftp timeout")
	giconfig.Add(Retry, 3, "ftp retry")
}
