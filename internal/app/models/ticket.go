package models

import "time"

type Ticket struct {
	Ticket_number  uint
	Flight_number  uint
	Place          string
	Price          uint
	Passenger_name string
	Terminal       uint
	Reg_time       time.Time
	Takeoff_time   time.Time
	Arrive_time    time.Time
}
