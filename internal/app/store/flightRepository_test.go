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
		FlightKey:          uint(123412),
		AirplaneKey:        1234123,
		StartingAirportKey: 2134,
		FinalAirportKey:    1234123,
		MaxTicketsNumber:   1234,
		SalesAgentKey:      12343,
		AirCarrierKey:      2134123,
		TakeoffTime:        time.Now(),
		ArriveTime:         time.Now(),
		Date:               time.Now(),
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
		FlightKey:          uint(key),
		AirplaneKey:        1234123,
		StartingAirportKey: 2134,
		FinalAirportKey:    1234123,
		MaxTicketsNumber:   1234,
		SalesAgentKey:      12343,
		AirCarrierKey:      2134123,
		TakeoffTime:        time.Now(),
		ArriveTime:         time.Now(),
		Date:               time.Now(),
	})

	flight, err := s.FlightRepository().FindByFlightKey(uint(key))

	assert.NoError(t, err)
	assert.NotNil(t, flight)
}
