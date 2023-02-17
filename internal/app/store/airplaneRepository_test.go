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
		AirplaneKey:   12312,
		NumberOfSits:  12,
		Model:         "Петух",
		AirCarrierKey: 12342134,
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
		AirplaneKey:   123122,
		NumberOfSits:  12,
		Model:         name,
		AirCarrierKey: 12342134,
	})

	airplane, err := s.AirplaneRepository().FindByModel(name)

	assert.NoError(t, err)
	assert.NotNil(t, airplane)
}
