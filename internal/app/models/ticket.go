package models

import "time"

type Ticket struct {
	TicketNumber  uint
	FlightNumber  uint
	Place         string
	Price         uint
	PassengerName string
	Terminal      uint
	RegTime       time.Time
	TakeoffTime   time.Time
	ArriveTime    time.Time
}
