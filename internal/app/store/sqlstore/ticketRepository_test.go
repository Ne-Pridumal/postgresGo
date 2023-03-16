package sqlstore_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store/sqlstore"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var bookRef = 1234

func TestTicketRepository(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("tickets")
	defer teardown("bookings")
	s := sqlstore.New(db)
	s.Bookings().Create(&models.Booking{
		BookRef:     bookRef,
		TotalAmount: 1234,
		BookDate:    time.Now(),
	})

	ticket, err := s.Tickets().Create(&models.Ticket{
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
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("tickets")
	defer teardown("bookings")
	s := sqlstore.New(db)
	name := "ИРИСКИН ЕГОР"

	_, err := s.Tickets().FindByName(name)
	assert.Error(t, err)

	s.Bookings().Create(&models.Booking{
		BookRef:     bookRef,
		TotalAmount: 1234,
		BookDate:    time.Now(),
	})
	s.Tickets().Create(&models.Ticket{
		BookRef:       bookRef,
		Key:           "123",
		PassengerId:   1234,
		ContactDate:   time.Now(),
		PassengerName: name,
	})

	ticket, err := s.Tickets().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
}
