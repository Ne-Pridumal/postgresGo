package models

import "time"

type Booking struct {
	BookRef     int       `json:"ref"`
	BookDate    time.Time `json:"date"`
	TotalAmount int       `json:"amount"`
}
