package apiserver

// Config...
type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel:    "debug",
		DatabaseURL: "host=localhost port=5432 user=root password=root dbname=pg_db sslmode=disable",
	}
}
