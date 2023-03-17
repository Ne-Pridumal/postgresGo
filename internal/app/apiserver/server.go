package apiserver

import (
	"ne-pridumal/go-postgress/internal/app/apiserver/handlers"
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router   *mux.Router
	logger   *logrus.Logger
	handlers *handlers.Handlers
}

func newServer(store store.Store) *server {
	s := &server{
		router:   mux.NewRouter(),
		logger:   logrus.New(),
		handlers: handlers.NewHandler(store),
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	//bookings route
	s.router.HandleFunc("/bookings", s.handlers.Bookings().Create()).Methods("POST")
	s.router.HandleFunc("/bookings", s.handlers.Bookings().GetAll()).Methods("GET")
}
