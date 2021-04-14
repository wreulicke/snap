package snapshot

import (
	"flag"
)

type Config struct {
	UpdateSnapshot bool
}

var (
	update = flag.Bool("update", false, "Update snapshot files")
)

type Option func(c Config) Config

func NewConfig(options ...Option) *Config {
	c := Config{
		UpdateSnapshot: *update,
	}
	for _, e := range options {
		c = e(c)
	}
	return &c
}
