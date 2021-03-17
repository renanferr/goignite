package jlaffaye

import "github.com/b2wdigital/goignite/v2/config"

const (
	root     = "gi.jlaffaye"
	addr     = root + ".addr"
	username = root + ".username"
	password = root + ".password"
	timeout  = root + ".timeout"
	retry    = root + ".retry"
)

func init() {
	config.Add(addr, "", "ftp address")
	config.Add(username, "", "ftp username")
	config.Add(password, "", "ftp password")
	config.Add(timeout, 10, "ftp timeout")
	config.Add(retry, 3, "ftp retry")
}
