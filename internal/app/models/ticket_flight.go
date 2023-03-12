package models

import (
	"time"
)

type TicketFlights struct {
	TicketKey          string
	FlightKey          int
	ScheduledDeparture time.Time
	FareCondition      string
	Amount             int
}
