package apiserver

import "ne-pridumal/go-postgress/internal/app/store"

// Config...
type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	Store       *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
		Store:       store.NewConfig(),
	}
}
