package models

import "time"

type Ticket struct {
	Key           string
	BookRef       int
	PassengerId   int
	PassengerName string
	ContactDate   time.Time
}
