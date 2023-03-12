package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

var aircraftKey = 1232143

func TestSeatRepository(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("seats")
	teardown("aircrafts")
	s := store.New(db)

	s.AircraftRepository().Create(&models.Aircraft{
		Key:   aircraftKey,
		Model: "asdfasdf",
		Range: 2134231,
	})

	seat, err := s.SeatRepository().Create(&models.Seat{
		AircraftKey:    aircraftKey,
		Number:         "asdf232",
		FareConditions: "sadfas",
	})

	assert.NoError(t, err)
	assert.NotNil(t, seat)
}
