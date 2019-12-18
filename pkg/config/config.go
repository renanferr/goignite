package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"path/filepath"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"

	flag "github.com/spf13/pflag"
)

var (
	entries  []config
	Instance *koanf.Koanf
	f        *flag.FlagSet
)

func init() {
	prepare()
}

type config struct {
	key         string
	example     interface{}
	description string
}

func Add(key string, example interface{}, description string) {
	entries = append(entries, config{
		key:         key,
		example:     example,
		description: description,
	})
}

func prepare() {

	entries = []config{}

	Instance = koanf.New(".")

	// Use the POSIX compliant pflag lib instead of Go's flag lib.
	f = flag.NewFlagSet("config", flag.ContinueOnError)

	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

}

func Parse() error {

	// Parse flags
	parseFlags()

	// Load the config files provided in the commandline.
	files, _ := f.GetStringSlice("conf")
	for _, c := range files {

		var parser koanf.Parser

		if filepath.Ext(c) == ".toml" {
			parser = toml.Parser()
		} else if filepath.Ext(c) == ".yaml" || filepath.Ext(c) == ".yml" {
			parser = yaml.Parser()
		} else if filepath.Ext(c) == ".json" {
			parser = json.Parser()
		} else {
			return errors.New(fmt.Sprintf("error on check extension of file %s", c))
		}

		if err := Instance.Load(file.Provider(c), parser); err != nil {
			return err
		}
	}

	// Env vars
	err := Instance.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "")), "_", ".", -1)
	}), nil)
	if err != nil {
		return err
	}

	// Load flags
	flap := posflag.Provider(f, ".", Instance)

	if err := Instance.Load(flap, nil); err != nil {
		return err
	}

	return nil
}

func Unmarshal(o interface{}) error {
	return Instance.UnmarshalWithConf("", &o, koanf.UnmarshalConf{Tag: "config"})
}

func parseFlags() {

	for _, v := range entries {

		switch t := v.example.(type) {

		case int:
			f.Int(v.key, t, v.description)
		case string:
			f.String(v.key, t, v.description)
		case bool:
			f.Bool(v.key, t, v.description)
		default:
			fmt.Println("type unknown")
		}

	}

	// Path to one or more config files to load into koanf along with some config params.
	f.StringSlice("conf", nil, "path to one or more config files")

	f.Parse(os.Args[0:])
}
