package store

type Store interface {
	Tickets() TicketRepository
	Aircrafts() AircraftRepository
	Bookings() BookingRepository
	Airports() AirportRepository
	Flights() FlightRepository
	Seats() SeatRepository
	TicketFlights() TicketFlightRepository
	BoardingPasses() BoardingPassRepository
}
