package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAirportAgentRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("airport_agents")

	airportAgent, err := s.AirportAgentRepository().Create(&models.AirportAgent{
		AirportKey:      123412,
		Name:            "sdaf",
		AirCarrierKey:   2134,
		AirportAgentKey: 2134,
	})

	assert.NoError(t, err)
	assert.NotNil(t, airportAgent)
}

func TestAirportAgentRepository_FindByName(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("airport_agents")

	name := "Аэропорт Колотушкино"

	_, err := s.AirportAgentRepository().FindByName(name)
	assert.Error(t, err)

	s.AirportAgentRepository().Create(&models.AirportAgent{
		AirportAgentKey: 1234,
		AirCarrierKey:   1234,
		Name:            name,
		AirportKey:      1234,
	})

	airportAgent, err := s.AirportAgentRepository().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, airportAgent)
}
