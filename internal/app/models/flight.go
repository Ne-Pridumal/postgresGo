package models

import "time"

type Flight struct {
	Flight_key           uint
	Airplane_key         uint
	Starting_airport_key uint
	Final_airport_key    uint
	Max_tickets_number   uint
	Sales_agent_key      uint
	Air_carrier_key      uint
	Takeoff_time         time.Time
	Arrive_time          time.Time
	Date                 time.Time
}
