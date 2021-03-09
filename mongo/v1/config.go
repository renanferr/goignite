package gimongo

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root     = "gi.mongo"
	uri      = root + ".uri"
	authRoot = root + ".auth"
	username = authRoot + ".username"
	password = authRoot + ".password"
	ExtRoot  = root + ".ext"
)

func init() {

	giconfig.Add(uri, "mongodb://localhost:27017/temp", "define mongodb uri")

	giconfig.Add(username, "", "define mongodb username")

	giconfig.Add(password, "", "define mongodb password")
}
