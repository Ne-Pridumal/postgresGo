package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTicketFlightRepository(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("ticket_flights")
	s := store.New(db)
	fkey := 6345
	tkey := "abc123"
	CreateFlightAndTickets(t, fkey, tkey)
	ticket, err := s.TicketFlightRepository().Create(&models.TicketFlights{
		FlightKey:     fkey,
		TicketKey:     tkey,
		Amount:        1223,
		FareCondition: "sdfsdf",
		ScheduledDeparture: time.Date(
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			12,
			12,
			12,
			12,
			time.Now().UTC().Location(),
		),
	})

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
}

func CreateFlightAndTickets(t *testing.T, fkey int, tkey string) {
	CreateAircraftsAndAirports(t)
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("flights")
	defer teardown("tickets")
	s := store.New(db)
	s.FlightRepository().Create(&models.Flight{
		FlightKey:        fkey,
		ActualArrival:    time.Now(),
		ActualDeparture:  time.Now(),
		AircraftKey:      key,
		ArrivalAirport:   key,
		Status:           "sdf",
		DepartureAirport: key,
		ScheduledArrival: time.Now(),
		ScheduledDeparture: time.Date(
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			12,
			12,
			12,
			12,
			time.Now().UTC().Location(),
		),
	})
	s.BookingRepository().Create(&models.Booking{
		BookRef:     bookRef,
		TotalAmount: 1234,
		BookDate:    time.Now(),
	})
	s.TicketRepository().Create(&models.Ticket{
		BookRef:       bookRef,
		Key:           tkey,
		PassengerId:   1234,
		ContactDate:   time.Now(),
		PassengerName: "asdfasdf",
	})
}
