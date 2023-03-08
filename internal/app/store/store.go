package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config *Config

	db                 *sql.DB
	bookingRepository  *BookingRepository
	ticketRepository   *TicketRepository
	airportRepository  *AirportRepository
	flightRepository   *FlightRepository
	aircraftRepository *AircraftRepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) BookingRepository() *BookingRepository {
	if s.bookingRepository != nil {
		return s.bookingRepository
	}
	s.bookingRepository = &BookingRepository{
		store: s,
	}

	return s.bookingRepository
}

func (s *Store) TicketRepository() *TicketRepository {
	if s.ticketRepository != nil {
		return s.ticketRepository
	}

	s.ticketRepository = &TicketRepository{
		store: s,
	}

	return s.ticketRepository
}

func (s *Store) AircraftRepository() *AircraftRepository {
	if s.aircraftRepository != nil {
		return s.aircraftRepository
	}

	s.aircraftRepository = &AircraftRepository{
		store: s,
	}

	return s.aircraftRepository
}

func (s *Store) AirportRepository() *AirportRepository {
	if s.airportRepository != nil {
		return s.airportRepository
	}

	s.airportRepository = &AirportRepository{
		store: s,
	}

	return s.airportRepository
}
