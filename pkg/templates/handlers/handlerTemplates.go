package handlers

import "net/http"

type templatesActions interface {
	Render() http.HandlerFunc
}

type HandlerTemplates struct {
	mainPage *mainPage
}

func NewHandlerTemplates() *HandlerTemplates {
	return &HandlerTemplates{}
}

func (h *HandlerTemplates) MainPage() templatesActions {
	if h.mainPage != nil {
		return h.mainPage
	}
	h.mainPage = &mainPage{}
	return h.mainPage
}
