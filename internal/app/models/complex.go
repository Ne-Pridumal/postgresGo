package models

import "time"

type Complex struct {
	Destinations []Airport
	Departures   []Airport
	Fligts       []Flight
	Tickets      []Ticket
}

type ComplexQuery struct {
	Desination   string
	Departure    string
	Date         time.Time
	TicketAmount int
}
