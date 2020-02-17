package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

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
	instance *koanf.Koanf
	f        *flag.FlagSet
)

func init() {
	prepare()
}

type config struct {
	key         string
	value       interface{}
	description string
}

func Add(key string, value interface{}, description string) {
	entries = append(entries, config{
		key:         key,
		value:       value,
		description: description,
	})
}

func prepare() {

	entries = []config{}

	instance = koanf.New(".")

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

		if err := instance.Load(file.Provider(c), parser); err != nil {
			return err
		}
	}

	// Env vars
	err := instance.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "")), "_", ".", -1)
	}), nil)
	if err != nil {
		return err
	}

	// Load flags
	flap := posflag.Provider(f, ".", instance)

	if err := instance.Load(flap, nil); err != nil {
		return err
	}

	return nil
}

func parseFlags() {

	for _, v := range entries {

		switch t := v.value.(type) {

		case string:
			f.String(v.key, t, v.description)
		case []string:
			f.StringSlice(v.key, t, v.description)
		case bool:
			f.Bool(v.key, t, v.description)
		case []bool:
			f.BoolSlice(v.key, t, v.description)
		case int:
			f.Int(v.key, t, v.description)
		case []int:
			f.IntSlice(v.key, t, v.description)
		case int64:
			f.Int64(v.key, t, v.description)
		case int32:
			f.Int32(v.key, t, v.description)
		case int16:
			f.Int16(v.key, t, v.description)
		case int8:
			f.Int8(v.key, t, v.description)
		case uint:
			f.Uint(v.key, t, v.description)
		case []uint:
			f.UintSlice(v.key, t, v.description)
		case uint64:
			f.Uint64(v.key, t, v.description)
		case uint32:
			f.Uint32(v.key, t, v.description)
		case uint16:
			f.Uint16(v.key, t, v.description)
		case uint8:
			f.Uint8(v.key, t, v.description)
		case time.Duration:
			f.Duration(v.key, t, v.description)
		case []time.Duration:
			f.DurationSlice(v.key, t, v.description)
		case []byte:
			f.BytesBase64(v.key, t, v.description)
		case float32:
			f.Float32(v.key, t, v.description)
		case float64:
			f.Float64(v.key, t, v.description)
		case net.IP:
			f.IP(v.key, t, v.description)
		case []net.IP:
			f.IPSlice(v.key, t, v.description)
		case net.IPMask:
			f.IPMask(v.key, t, v.description)
		default:
			fmt.Println("type unknown")
		}

	}

	// Path to one or more config files to load into koanf along with some config params.
	f.StringSlice("conf", nil, "path to one or more config files")

	f.Parse(os.Args[0:])
}
