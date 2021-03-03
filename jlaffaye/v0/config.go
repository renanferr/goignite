package gijlaffaye

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root     = "gi.jlaffaye"
	addr     = root + ".addr"
	username = root + ".username"
	password = root + ".password"
	timeout  = root + ".timeout"
	retry    = root + ".retry"
)

func init() {
	giconfig.Add(addr, "", "ftp address")
	giconfig.Add(username, "", "ftp username")
	giconfig.Add(password, "", "ftp password")
	giconfig.Add(timeout, 10, "ftp timeout")
	giconfig.Add(retry, 3, "ftp retry")
}
