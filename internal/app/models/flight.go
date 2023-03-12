package models

import (
	"time"
)

type Flight struct {
	FlightKey          int
	ScheduledDeparture time.Time
	ScheduledArrival   time.Time
	DepartureAirport   int
	ArrivalAirport     int
	Status             string
	AircraftKey        int
	ActualDeparture    time.Time
	ActualArrival      time.Time
}
