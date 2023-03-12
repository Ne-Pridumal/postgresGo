package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db                     *sql.DB
	bookingRepository      *BookingRepository
	ticketRepository       *TicketRepository
	airportRepository      *AirportRepository
	flightRepository       *FlightRepository
	aircraftRepository     *AircraftRepository
	seatRepository         *SeatRepository
	ticketFlightRepository *TicketFlightRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) BookingRepository() *BookingRepository {
	if s.bookingRepository != nil {
		return s.bookingRepository
	}
	s.bookingRepository = &BookingRepository{
		store: s.db,
	}

	return s.bookingRepository
}

func (s *Store) TicketRepository() *TicketRepository {
	if s.ticketRepository != nil {
		return s.ticketRepository
	}

	s.ticketRepository = &TicketRepository{
		store: s.db,
	}

	return s.ticketRepository
}

func (s *Store) AircraftRepository() *AircraftRepository {
	if s.aircraftRepository != nil {
		return s.aircraftRepository
	}

	s.aircraftRepository = &AircraftRepository{
		store: s.db,
	}

	return s.aircraftRepository
}

func (s *Store) AirportRepository() *AirportRepository {
	if s.airportRepository != nil {
		return s.airportRepository
	}

	s.airportRepository = &AirportRepository{
		store: s.db,
	}

	return s.airportRepository
}

func (s *Store) SeatRepository() *SeatRepository {
	if s.seatRepository != nil {
		return s.seatRepository
	}

	s.seatRepository = &SeatRepository{
		store: s.db,
	}
	return s.seatRepository
}

func (s *Store) FlightRepository() *FlightRepository {
	if s.flightRepository != nil {
		return s.flightRepository
	}

	s.flightRepository = &FlightRepository{
		store: s.db,
	}
	return s.flightRepository
}

func (s *Store) TicketFlightRepository() *TicketFlightRepository {
	if s.ticketFlightRepository != nil {
		return s.ticketFlightRepository
	}

	s.ticketFlightRepository = &TicketFlightRepository{
		store: s.db,
	}

	return s.ticketFlightRepository
}
