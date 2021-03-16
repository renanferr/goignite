package main

import (
	"os"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
)

type AppConfig struct {
	Application struct {
		Name       string
		MyNameTest string
	}
}

func init() {
	giconfig.Add("app.application.name", "app_test", "name of application")
	giconfig.Add("app.application.myName", "my_name_test", "name of application")
}

func main() {

	os.Setenv("APP_APPLICATION_NAME", "app_test_env")
	os.Setenv("APP_APPLICATION_MY.NAME.TEST", "my_name_test_env")
	os.Setenv("CONF", "./examples/config/config.yaml")

	giconfig.Load()
	gilogrus.NewLogger()

	c := AppConfig{}

	giconfig.UnmarshalWithPath("app", &c)

	gilog.Info(c.Application.Name)
	gilog.Info(c.Application.MyNameTest)
}
