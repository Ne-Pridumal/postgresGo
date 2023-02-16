package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFlightRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("flights")

	ticket, err := s.FlightRepository().Create(&models.Flight{
		Flight_key:           uint(123412),
		Airplane_key:         1234123,
		Starting_airport_key: 2134,
		Final_airport_key:    1234123,
		Max_tickets_number:   1234,
		Sales_agent_key:      12343,
		Air_carrier_key:      2134123,
		Takeoff_time:         time.Now(),
		Arrive_time:          time.Now(),
		Date:                 time.Now(),
	})

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
}

func TestFlightRepository_FindByFlightKey(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("tickets")

	key := 12342134

	_, err := s.FlightRepository().FindByFlightKey(uint(key))
	assert.Error(t, err)

	s.FlightRepository().Create(&models.Flight{
		Flight_key:           uint(key),
		Airplane_key:         1234123,
		Starting_airport_key: 2134,
		Final_airport_key:    1234123,
		Max_tickets_number:   1234,
		Sales_agent_key:      12343,
		Air_carrier_key:      2134123,
		Takeoff_time:         time.Now(),
		Arrive_time:          time.Now(),
		Date:                 time.Now(),
	})

	flight, err := s.FlightRepository().FindByFlightKey(uint(key))

	assert.NoError(t, err)
	assert.NotNil(t, flight)
}
