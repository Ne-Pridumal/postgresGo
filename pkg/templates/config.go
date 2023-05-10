package templates

type Config struct {
	BindTemplateAddress string `toml:"bind_template_address"`
}

func NewConfig() *Config {
	return &Config{
		BindTemplateAddress: ":8181",
	}
}
