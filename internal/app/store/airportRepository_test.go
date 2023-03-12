package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAirportRepository(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("airports")
	s := store.New(db)

	airport, err := s.AirportRepository().Create(&models.Airport{
		Key:         1234,
		Name:        "sarwera",
		City:        "sadfasdf",
		Coordinates: "12341234233",
	})

	assert.NoError(t, err)
	assert.NotNil(t, airport)
}

func TestAirportRepository_FindByName(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("airports")
	s := store.New(db)
	name := "test"

	s.AirportRepository().Create(&models.Airport{
		Key:         1234,
		Name:        name,
		City:        "sadfasdf",
		Coordinates: "12341234233",
	})

	airport, err := s.AirportRepository().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, airport)
}
