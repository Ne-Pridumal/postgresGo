package sqlstore_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

var aircraftKey = 1232143

func TestSeatRepository(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("seats")
	teardown("aircrafts")
	s := sqlstore.New(db)

	s.Aircrafts().Create(&models.Aircraft{
		Key:   aircraftKey,
		Model: "asdfasdf",
		Range: 2134231,
	})

	seat, err := s.Seats().Create(&models.Seat{
		AircraftKey:    aircraftKey,
		Number:         "asdf232",
		FareConditions: "sadfas",
	})

	assert.NoError(t, err)
	assert.NotNil(t, seat)
}
