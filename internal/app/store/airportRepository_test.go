package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAirportRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("airports")

	airport, err := s.AirportRepository().Create(&models.Airport{
		Airport_key: 123,
		Name:        "аэропорт колотушкино",
		Address:     "колотушкино",
	})

	assert.NoError(t, err)
	assert.NotNil(t, airport)
}

func TestAirportRepository_FindByName(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("airports")

	name := "Аэропорт Колотушкино"

	_, err := s.AirportRepository().FindByName(name)
	assert.Error(t, err)

	s.AirportRepository().Create(&models.Airport{
		Airport_key: 123,
		Address:     "qwerwe",
		Name:        name,
	})

	airport, err := s.AirportRepository().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, airport)
}
