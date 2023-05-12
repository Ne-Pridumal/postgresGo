package handlers

import (
	"net/http"
)

type mainPage struct {
}

type data struct {
	Title   string
	Message string
}

func (p *mainPage) Render() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagesTemplates.Render(w, "main.html", &data{
			Title:   "asdf",
			Message: "caxzcx",
		})
	}
}
