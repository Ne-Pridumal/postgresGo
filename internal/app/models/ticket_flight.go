package models

import (
	"time"
)

type TicketFlights struct {
	TicketKey          string
	FlightKey          uint
	ScheduledDeparture time.Time
	FareCondition      string
	Amount             uint
}
