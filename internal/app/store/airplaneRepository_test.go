package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAirplaneRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)

	defer teardown("airplanes")

	airplane, err := s.AirplaneRepository().Create(&models.Airplane{
		Airplane_key:    12312,
		Number_of_sits:  12,
		Model:           "Петух",
		Air_carrier_key: 12342134,
	})

	assert.NoError(t, err)

	assert.NotNil(t, airplane)
}

func TestAirplaneRepository_FindByModel(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("airplanes")

	name := "test"

	_, err := s.AirplaneRepository().FindByModel(name)
	assert.Error(t, err)

	s.AirplaneRepository().Create(&models.Airplane{
		Airplane_key:    123122,
		Number_of_sits:  12,
		Model:           name,
		Air_carrier_key: 12342134,
	})

	airplane, err := s.AirportRepository().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, airplane)
}
