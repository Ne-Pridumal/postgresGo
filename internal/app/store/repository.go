package store

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"time"
)

type Model interface {
	models.Seat | models.Flight | models.Airport | models.Aircraft | models.Booking | models.Ticket | models.BoardingPass | models.TicketFlights
}

type Repository[T Model] interface {
	GetAll() ([]T, error)
	Create(*T) (*T, error)
}

type TicketRepository interface {
	Repository[models.Ticket]
	Delete(string, int) error
}

type AircraftRepository interface {
	Repository[models.Aircraft]
	Delete(int) error
}

type BookingRepository interface {
	Repository[models.Booking]
	Delete(int) error
}

type AirportRepository interface {
	Repository[models.Airport]
	Delete(int) error
}

type FlightRepository interface {
	Repository[models.Flight]
	Delete(int, time.Time) error
}

type SeatRepository interface {
	Repository[models.Seat]
	Delete(int, string) error
}

type TicketFlightRepository interface {
	Repository[models.TicketFlights]
	Delete(int, string, time.Time) error
}

type BoardingPassRepository interface {
	Repository[models.BoardingPass]
	Delete(int, string, time.Time) error
}
