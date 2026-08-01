package config

import (
	"flag"
)

type Config struct {
	PORT *string
}

func Load_config() Config {
	cfg := Config{}

	flag.StringVar(cfg.PORT, "PORT", ":4000", "HTTP network address")

	return cfg
}
