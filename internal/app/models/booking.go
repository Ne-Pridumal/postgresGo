package models

import "time"

type Booking struct {
	BookRef     int32
	BookDate    time.Time
	TotalAmount int32
}
