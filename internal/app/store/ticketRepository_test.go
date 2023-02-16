package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTicketRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("tickets")

	ticket, err := s.TicketRepository().Create(&models.Ticket{
		Ticket_number:  232,
		Flight_number:  123,
		Place:          "32a",
		Price:          32324,
		Terminal:       3,
		Arrive_time:    time.Now(),
		Takeoff_time:   time.Now(),
		Reg_time:       time.Now(),
		Passenger_name: "Борис",
	})

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
}

func TestTicketRepository_FindByName(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("tickets")

	name := "ИРИСКИН ЕГОР"

	_, err := s.TicketRepository().FindByName(name)
	assert.Error(t, err)

	s.TicketRepository().Create(&models.Ticket{
		Ticket_number:  232,
		Flight_number:  123,
		Place:          "32a",
		Price:          32324,
		Terminal:       3,
		Arrive_time:    time.Now(),
		Takeoff_time:   time.Now(),
		Reg_time:       time.Now(),
		Passenger_name: name,
	})

	ticket, err := s.TicketRepository().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
}
