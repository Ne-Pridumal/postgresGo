package models

import "time"

type Booking struct {
	BookRef     int
	BookDate    time.Time
	TotalAmount int
}
