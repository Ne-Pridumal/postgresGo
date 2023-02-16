package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config *Config

	db                     *sql.DB
	ticketRepository       *TicketRepository
	airportRepository      *AirportRepository
	airplaneRepository     *AirplaneRepository
	flightRepository       *FlightRepository
	salesAgentRepository   *SalesAgentRepository
	airCarrierRepository   *AirCarrierRepository
	airportAgentRepository *AirportAgentRepository
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

func (s *Store) TicketRepository() *TicketRepository {
	if s.ticketRepository != nil {
		return s.ticketRepository
	}

	s.ticketRepository = &TicketRepository{
		store: s,
	}

	return s.ticketRepository
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

func (s *Store) AirplaneRepository() *AirplaneRepository {
	if s.airplaneRepository != nil {
		return s.airplaneRepository
	}

	s.airplaneRepository = &AirplaneRepository{
		store: s,
	}

	return s.airplaneRepository
}

func (s *Store) FlightRepository() *FlightRepository {
	if s.flightRepository != nil {
		return s.flightRepository
	}

	s.flightRepository = &FlightRepository{
		store: s,
	}

	return s.flightRepository
}

func (s *Store) SalesAgentRepository() *SalesAgentRepository {
	if s.salesAgentRepository != nil {
		return s.salesAgentRepository
	}

	s.salesAgentRepository = &SalesAgentRepository{
		store: s,
	}

	return s.salesAgentRepository
}

func (s *Store) AirCarrierRepository() *AirCarrierRepository {
	if s.airCarrierRepository != nil {
		return s.airCarrierRepository
	}

	s.airCarrierRepository = &AirCarrierRepository{
		store: s,
	}

	return s.airCarrierRepository
}
func (s *Store) AirportAgentRepository() *AirportAgentRepository {
	if s.airportAgentRepository != nil {
		return s.airportAgentRepository
	}

	s.airportAgentRepository = &AirportAgentRepository{
		store: s,
	}

	return s.airportAgentRepository
}
