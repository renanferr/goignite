package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPFlag(t *testing.T) {
	prepare()
	Add("key", "value", "test")
	Parse()
	assert.Equal(t, "value", Instance.String("key"), "they should be equal")

}

func TestEnv(t *testing.T) {

	prepare()
	os.Setenv("K_ENV", "value")
	Parse()
	assert.Equal(t, "value", Instance.String("k.env"), "they should be equal")

}

func TestConf(t *testing.T) {

	os.Args = []string{"--conf", "./testdata/config.json", "--conf", "./testdata/config.yaml"}
	prepare()
	Parse()

	assert.True(t, Instance.Bool("debug"), "they should be true")
	assert.Equal(t, "127.0.0.13", Instance.String("redis.host"), "they should be equal")

}
