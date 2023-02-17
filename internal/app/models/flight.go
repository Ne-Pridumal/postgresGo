package models

import "time"

type Flight struct {
	FlightKey          uint
	AirplaneKey        uint
	StartingAirportKey uint
	FinalAirportKey    uint
	MaxTicketsNumber   uint
	SalesAgentKey      uint
	AirCarrierKey      uint
	TakeoffTime        time.Time
	ArriveTime         time.Time
	Date               time.Time
}
