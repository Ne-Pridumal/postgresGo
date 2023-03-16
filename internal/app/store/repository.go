package store

import "ne-pridumal/go-postgress/internal/app/models"

type Model interface {
	models.Seat | models.Flight | models.Airport | models.Aircraft | models.Booking | models.Ticket | models.BoardingPass | models.TicketFlights
}

type Repository[T Model] interface {
	DeleteById(int) error
	GetAll() ([]T, error)
	Create(*T) (*T, error)
}

type TicketRepository interface {
	Repository[models.Ticket]
	FindByName(string) (*models.Ticket, error)
}

type AircraftRepository interface {
	Repository[models.Aircraft]
}

type BookingRepository interface {
	Repository[models.Booking]
}

type AirportRepository interface {
	Repository[models.Airport]
}

type FlightRepository interface {
	Repository[models.Flight]
}

type SeatRepository interface {
	Repository[models.Seat]
}

type TicketFlightRepository interface {
	Repository[models.TicketFlights]
}
