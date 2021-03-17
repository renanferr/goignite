package mongo

import "github.com/b2wdigital/goignite/v2/config"

const (
	root     = "gi.mongo"
	uri      = root + ".uri"
	authRoot = root + ".auth"
	username = authRoot + ".username"
	password = authRoot + ".password"
	ExtRoot  = root + ".ext"
)

func init() {

	config.Add(uri, "mongodb://localhost:27017/temp", "define mongodb uri")

	config.Add(username, "", "define mongodb username")

	config.Add(password, "", "define mongodb password")
}
