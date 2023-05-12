package templates

import "net/http"

type Template struct {
}

func Start(config *Config) error {
	srv := newServer()
	return http.ListenAndServe(config.BindTemplateAddress, srv)
}
