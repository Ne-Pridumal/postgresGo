package sqlstore_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAirportRepository(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("airports")
	s := sqlstore.New(db)

	airport, err := s.Airports().Create(&models.Airport{
		Key:         1234,
		Name:        "sarwera",
		City:        "sadfasdf",
		Coordinates: "12341234233",
	})

	assert.NoError(t, err)
	assert.NotNil(t, airport)
}

func TestAirportRepository_FindByName(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("airports")
	s := sqlstore.New(db)
	name := "test"

	s.Airports().Create(&models.Airport{
		Key:         1234,
		Name:        name,
		City:        "sadfasdf",
		Coordinates: "12341234233",
	})

	airport, err := s.Airports().FindByName(name)

	assert.NoError(t, err)
	assert.NotNil(t, airport)
}
