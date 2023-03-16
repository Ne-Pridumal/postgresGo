package sqlstore_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store/sqlstore"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var key = 234123

func TestFlightRepository(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	CreateAircraftsAndAirports(t)
	defer teardown("flights")
	s := sqlstore.New(db)
	flight, err := s.Flights().Create(&models.Flight{
		FlightKey:          1234213,
		ActualArrival:      time.Now(),
		ActualDeparture:    time.Now(),
		AircraftKey:        key,
		ArrivalAirport:     key,
		Status:             "sdf",
		DepartureAirport:   key,
		ScheduledArrival:   time.Now(),
		ScheduledDeparture: time.Now(),
	})

	assert.NoError(t, err)
	assert.NotNil(t, flight)
}

func TestFlightRepository_FindbyName(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	CreateAircraftsAndAirports(t)
	defer teardown("flights")
	s := sqlstore.New(db)

	fKey := 423534234
	s.Flights().Create(&models.Flight{
		FlightKey:          fKey,
		ActualArrival:      time.Now(),
		ActualDeparture:    time.Now(),
		AircraftKey:        key,
		ArrivalAirport:     key,
		Status:             "sdf",
		DepartureAirport:   key,
		ScheduledArrival:   time.Now(),
		ScheduledDeparture: time.Now(),
	})

	flight, err := s.Flights().FindByFlightKey(fKey)

	assert.NoError(t, err)
	assert.NotNil(t, flight)
}

func CreateAircraftsAndAirports(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("airports")
	defer teardown("aircrafts")
	s := sqlstore.New(db)

	s.Airports().Create(&models.Airport{
		Key:         key,
		Coordinates: "2134213",
		City:        "asdfasdf",
		Name:        "asdfasdfsadf",
	})
	s.Aircrafts().Create(&models.Aircraft{
		Key:   key,
		Range: 134123,
		Model: "sadfsadf",
	})
}
