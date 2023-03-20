package sqlstore

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/store"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	db                     *sql.DB
	bookingRepository      *bookingRepository
	ticketRepository       *ticketRepository
	airportRepository      *airportRepository
	flightRepository       *flightRepository
	aircraftRepository     *aircraftRepository
	seatRepository         *seatRepository
	ticketFlightRepository *ticketFlightRepository
	boardingPassRepository *boardingPassRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Bookings() store.BookingRepository {
	if s.bookingRepository != nil {
		return s.bookingRepository
	}
	s.bookingRepository = &bookingRepository{
		store: s.db,
	}

	return s.bookingRepository
}

func (s *Store) Tickets() store.TicketRepository {
	if s.ticketRepository != nil {
		return s.ticketRepository
	}

	s.ticketRepository = &ticketRepository{
		store: s.db,
	}

	return s.ticketRepository
}

func (s *Store) Aircrafts() store.AircraftRepository {
	if s.aircraftRepository != nil {
		return s.aircraftRepository
	}

	s.aircraftRepository = &aircraftRepository{
		store: s.db,
	}

	return s.aircraftRepository
}

func (s *Store) Airports() store.AirportRepository {
	if s.airportRepository != nil {
		return s.airportRepository
	}

	s.airportRepository = &airportRepository{
		store: s.db,
	}

	return s.airportRepository
}

func (s *Store) Seats() store.SeatRepository {
	if s.seatRepository != nil {
		return s.seatRepository
	}

	s.seatRepository = &seatRepository{
		store: s.db,
	}
	return s.seatRepository
}

func (s *Store) Flights() store.FlightRepository {
	if s.flightRepository != nil {
		return s.flightRepository
	}

	s.flightRepository = &flightRepository{
		store: s.db,
	}
	return s.flightRepository
}

func (s *Store) TicketFlights() store.TicketFlightRepository {
	if s.ticketFlightRepository != nil {
		return s.ticketFlightRepository
	}

	s.ticketFlightRepository = &ticketFlightRepository{
		store: s.db,
	}

	return s.ticketFlightRepository
}

func (s *Store) BoardingPasses() store.BoardingPassRepository {
	if s.boardingPassRepository != nil {
		return s.boardingPassRepository
	}

	s.boardingPassRepository = &boardingPassRepository{
		store: s.db,
	}

	return s.boardingPassRepository

}
