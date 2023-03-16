package store

type Store interface {
	Tickets() TicketRepository
	Aircrafts() AircraftRepository
	Bookings() BookingRepository
	Airports() AircraftRepository
	Flights() FlightRepository
	Seats() SeatRepository
	TicketFlights() TicketFlightRepository
}
