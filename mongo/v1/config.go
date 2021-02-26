package gimongo

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ConfigRoot = "gi.mongo"
	Uri        = ConfigRoot + ".uri"
	AuthRoot   = ConfigRoot + ".auth"
	Username   = AuthRoot + ".username"
	Password   = AuthRoot + ".password"
)

func init() {

	giconfig.Add(Uri, "mongodb://localhost:27017/temp", "define mongodb uri")

	giconfig.Add(Username, "", "define username")

	giconfig.Add(Password, "", "define password")
}
