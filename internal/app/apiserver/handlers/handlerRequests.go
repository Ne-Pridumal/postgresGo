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

type HandlerRequests struct {
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

func NewHandler(store store.Store) *HandlerRequests {
	return &HandlerRequests{
		store: store,
	}
}

func (h *HandlerRequests) Bookings() actions {
	if h.bookings != nil {
		return h.bookings
	}
	h.bookings = &bookingActions{
		Store: h.store,
	}
	return h.bookings
}

func (h *HandlerRequests) Tickets() actions {
	if h.tickets != nil {
		return h.tickets
	}
	h.tickets = &ticketsActions{
		Store: h.store,
	}
	return h.tickets
}

func (h *HandlerRequests) Aircrafts() actions {
	if h.aircrafts != nil {
		return h.aircrafts
	}
	h.aircrafts = &aircraftActions{
		Store: h.store,
	}
	return h.aircrafts
}

func (h *HandlerRequests) Seats() actions {
	if h.seats != nil {
		return h.seats
	}
	h.seats = &seatsActions{
		Store: h.store,
	}
	return h.seats
}

func (h *HandlerRequests) Airports() actions {
	if h.airports != nil {
		return h.airports
	}
	h.airports = &airportActions{
		Store: h.store,
	}
	return h.airports
}

func (h *HandlerRequests) Flights() actions {
	if h.flights != nil {
		return h.flights
	}
	h.flights = &flightActions{
		Store: h.store,
	}
	return h.flights
}

func (h *HandlerRequests) TicketsFlights() actions {
	if h.ticketFlights != nil {
		return h.ticketFlights
	}
	h.ticketFlights = &ticketFlightsActions{
		Store: h.store,
	}
	return h.ticketFlights
}

func (h *HandlerRequests) BoardingPasses() actions {
	if h.boardingPasses != nil {
		return h.boardingPasses
	}
	h.boardingPasses = &boardingPassesActions{
		Store: h.store,
	}
	return h.boardingPasses

}
