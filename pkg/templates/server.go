package templates

import (
	"ne-pridumal/go-postgress/pkg/templates/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router   *mux.Router
	handlers *handlers.HandlerTemplates
}

func newServer() *server {
	s := &server{
		router:   mux.NewRouter(),
		handlers: handlers.NewHandlerTemplates(),
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handlers.MainPage().Render())
}
