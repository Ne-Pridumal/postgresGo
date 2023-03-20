package handlers

import (
	"ne-pridumal/go-postgress/internal/app/store"
	"net/http"
)

type actions interface {
	GetAll() http.HandlerFunc
	Create() http.HandlerFunc
	Delete() http.HandlerFunc
}

type Handlers struct {
	store          store.Store
	bookings       *bookingActions
	tickets        *ticketsActions
	aircrafts      *aircraftActions
	seats          *seatsActions
	airports       *airportActions
	flights        *flightActions
	ticketFlights  *ticketFlightsActions
	boardingPasses *boardingPassesActions
}

func NewHandler(store store.Store) *Handlers {
	return &Handlers{
		store: store,
	}
}

func (h *Handlers) Bookings() actions {
	if h.bookings != nil {
		return h.bookings
	}
	h.bookings = &bookingActions{
		Store: h.store,
	}
	return h.bookings
}

func (h *Handlers) Tickets() actions {
	if h.tickets != nil {
		return h.tickets
	}
	h.tickets = &ticketsActions{
		Store: h.store,
	}
	return h.tickets
}

func (h *Handlers) Aircrafts() actions {
	if h.aircrafts != nil {
		return h.aircrafts
	}
	h.aircrafts = &aircraftActions{
		Store: h.store,
	}
	return h.aircrafts
}

func (h *Handlers) Seats() actions {
	if h.seats != nil {
		return h.seats
	}
	h.seats = &seatsActions{
		Store: h.store,
	}
	return h.seats
}

func (h *Handlers) Airports() actions {
	if h.airports != nil {
		return h.airports
	}
	h.airports = &airportActions{
		Store: h.store,
	}
	return h.airports
}

func (h *Handlers) Flights() actions {
	if h.flights != nil {
		return h.flights
	}
	h.flights = &flightActions{
		Store: h.store,
	}
	return h.flights
}

func (h *Handlers) TicketsFlights() actions {
	if h.ticketFlights != nil {
		return h.ticketFlights
	}
	h.ticketFlights = &ticketFlightsActions{
		Store: h.store,
	}
	return h.ticketFlights
}

func (h *Handlers) BoardingPasses() actions {
	if h.boardingPasses != nil {
		return h.boardingPasses
	}
	h.boardingPasses = &boardingPassesActions{
		Store: h.store,
	}
	return h.boardingPasses

}
