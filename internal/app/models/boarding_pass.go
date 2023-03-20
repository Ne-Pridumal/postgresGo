package models

import "time"

type BoardingPass struct {
	TicketKey          string
	FlightKey          int
	ScheduledDeparture time.Time
	BoardingNo         string
	SeatNo             string
}
