package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DBConfig struct {
	Username string `config:"username"`
}

type Config struct {
	Addr      string
	Debug     string
	DB        DBConfig `config:"db"`
	Redis     struct {
		Host string `config:"h"`
	} `config:"red"`
}

func TestPFlag(t *testing.T) {

	prepare()

	Add("key", "value", "test")

	Parse()

	assert.Equal(t, "value", Instance.String("key"), "they should be equal")
}

func TestEnv(t *testing.T) {

	prepare()

	os.Setenv("K_ENV", "value")

	err := Parse()
	assert.Nil(t, err, "they should be nil")

	assert.Equal(t, "value", Instance.String("k.env"), "they should be equal")
}

func TestConf(t *testing.T) {

	prepare()

	os.Args = []string{"--conf", "./testdata/config.json", "--conf", "./testdata/config.yaml"}

	err := Parse()
	assert.Nil(t, err, "they should be nil")

	assert.True(t, Instance.Bool("debug"), "they should be true")
	assert.Equal(t, "127.0.0.13", Instance.String("redis.host"), "they should be equal")
}

func TestUnmarshal(t *testing.T) {

	prepare()

	var err error

	os.Args = []string{"--conf", "./testdata/config.json", "--conf", "./testdata/config.yaml"}

	err = Parse()
	assert.Nil(t, err, "they should be nil")

	c := Config{}
	err = Unmarshal(&c)
	assert.Nil(t, err, "they should be nil")
	assert.Equal(t, ":8083", c.Addr, "they should be equal")
	assert.Equal(t, "foosss", c.DB.Username, "they should be equal")
	assert.Equal(t, "127.0.0.14", c.Redis.Host, "they should be equal")
}
