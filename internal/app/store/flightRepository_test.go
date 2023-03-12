package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var key = 234123

func TestFlightRepository(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	CreateAircraftsAndAirports(t)
	defer teardown("flights")
	s := store.New(db)
	flight, err := s.FlightRepository().Create(&models.Flight{
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
	db, teardown := store.TestDB(t, databaseURL)
	CreateAircraftsAndAirports(t)
	defer teardown("flights")
	s := store.New(db)

	fKey := 423534234
	s.FlightRepository().Create(&models.Flight{
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

	flight, err := s.FlightRepository().FindByFlightKey(fKey)

	assert.NoError(t, err)
	assert.NotNil(t, flight)
}

func CreateAircraftsAndAirports(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("airports")
	defer teardown("aircrafts")
	s := store.New(db)

	s.AirportRepository().Create(&models.Airport{
		Key:         key,
		Coordinates: "2134213",
		City:        "asdfasdf",
		Name:        "asdfasdfsadf",
	})
	s.AircraftRepository().Create(&models.Aircraft{
		Key:   key,
		Range: 134123,
		Model: "sadfsadf",
	})
}
