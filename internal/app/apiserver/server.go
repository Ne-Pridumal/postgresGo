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
	s.configurateBookings()
	s.configurateTickets()
	s.configurateAircrafts()
	s.configurateSeats()
	s.configurateAirports()
	s.configurateFlights()
	s.configurateTicketFlights()
	s.configurateBoardingPasses()
}

func (s *server) configurateBookings() {
	s.router.HandleFunc("/bookings", s.handlers.Bookings().Create()).Methods("POST")
	s.router.HandleFunc("/bookings", s.handlers.Bookings().GetAll()).Methods("GET")
	s.router.HandleFunc("/bookings", s.handlers.Bookings().Delete()).Methods("DELETE")
}

func (s *server) configurateTickets() {
	s.router.HandleFunc("/tickets", s.handlers.Tickets().Create()).Methods("POST")
	s.router.HandleFunc("/tickets", s.handlers.Tickets().GetAll()).Methods("GET")
	s.router.HandleFunc("/tickets", s.handlers.Tickets().Delete()).Methods("DELETE")
}

func (s *server) configurateAircrafts() {
	s.router.HandleFunc("/aircrafts", s.handlers.Aircrafts().Create()).Methods("POST")
	s.router.HandleFunc("/aircrafts", s.handlers.Aircrafts().GetAll()).Methods("GET")
	s.router.HandleFunc("/aircrafts", s.handlers.Aircrafts().Delete()).Methods("DELETE")
}

func (s *server) configurateSeats() {
	s.router.HandleFunc("/seats", s.handlers.Seats().Create()).Methods("POST")
	s.router.HandleFunc("/seats", s.handlers.Seats().GetAll()).Methods("GET")
	s.router.HandleFunc("/seats", s.handlers.Seats().Delete()).Methods("DELETE")
}

func (s *server) configurateAirports() {
	s.router.HandleFunc("/airports", s.handlers.Airports().Create()).Methods("POST")
	s.router.HandleFunc("/airports", s.handlers.Airports().GetAll()).Methods("GET")
	s.router.HandleFunc("/airports", s.handlers.Airports().Delete()).Methods("DELETE")
}

func (s *server) configurateFlights() {
	s.router.HandleFunc("/flights", s.handlers.Flights().Create()).Methods("POST")
	s.router.HandleFunc("/flights", s.handlers.Flights().GetAll()).Methods("GET")
	s.router.HandleFunc("/flights", s.handlers.Flights().Delete()).Methods("DELETE")
}

func (s *server) configurateTicketFlights() {
	s.router.HandleFunc("/ticketFlights", s.handlers.TicketsFlights().Create()).Methods("POST")
	s.router.HandleFunc("/ticketFlights", s.handlers.TicketsFlights().GetAll()).Methods("GET")
	s.router.HandleFunc("/ticketFlights", s.handlers.TicketsFlights().Delete()).Methods("DELETE")
}

func (s *server) configurateBoardingPasses() {
	s.router.HandleFunc("/boardingPasses", s.handlers.BoardingPasses().Create()).Methods("POST")
	s.router.HandleFunc("/boardingPasses", s.handlers.BoardingPasses().GetAll()).Methods("GET")
	s.router.HandleFunc("/boardingPasses", s.handlers.BoardingPasses().Delete()).Methods("DELETE")
}
