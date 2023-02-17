package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAirCarrierRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("air_carriers")

	airCarrier, err := s.AirCarrierRepository().Create(&models.AirCarrier{
		Name:          "1234",
		AirCarrierKey: 2134,
	})

	assert.NoError(t, err)
	assert.NotNil(t, airCarrier)
}

func TestAirCarrierRepository_FindByName(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("air_carriers")

	name := "airCarriers"

	_, err := s.AirCarrierRepository().FindByName(name)
	assert.Error(t, err)

	s.AirCarrierRepository().Create(&models.AirCarrier{
		AirCarrierKey: 2134,
		Name:          name,
	})

	airCarrier, err := s.AirCarrierRepository().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, airCarrier)
}
