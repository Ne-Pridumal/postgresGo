package models

import "time"

type BoardingPass struct {
	TicketKey          string
	FlightKey          uint
	ScheduledDeparture time.Time
	BoardingNo         string
	SeatNo             string
}
