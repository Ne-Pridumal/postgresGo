package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var bookRef = 1234

func TestTicketRepository(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("tickets")
	teardown("bookings")
	s := store.New(db)
	s.BookingRepository().Create(&models.Booking{
		BookRef:     bookRef,
		TotalAmount: 1234,
		BookDate:    time.Now(),
	})

	ticket, err := s.TicketRepository().Create(&models.Ticket{
		BookRef:       bookRef,
		Key:           "123",
		PassengerId:   1234,
		ContactDate:   time.Now(),
		PassengerName: "Борис",
	})

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
}

func TestTicketRepository_FindByName(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("tickets")
	teardown("bookings")
	s := store.New(db)
	name := "ИРИСКИН ЕГОР"

	_, err := s.TicketRepository().FindByName(name)
	assert.Error(t, err)

	s.BookingRepository().Create(&models.Booking{
		BookRef:     bookRef,
		TotalAmount: 1234,
		BookDate:    time.Now(),
	})
	s.TicketRepository().Create(&models.Ticket{
		BookRef:       bookRef,
		Key:           "123",
		PassengerId:   1234,
		ContactDate:   time.Now(),
		PassengerName: name,
	})

	ticket, err := s.TicketRepository().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
}
