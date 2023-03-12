package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAircraftRepository(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("aircrafts")
	s := store.New(db)
	aircraft, err := s.AircraftRepository().Create(&models.Aircraft{
		Key:   1234,
		Model: "123421342",
		Range: 234324,
	})

	assert.NoError(t, err)
	assert.NotNil(t, aircraft)
}

func TestAircraftRepository_FindByModel(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("aircrafts")
	s := store.New(db)
	model := "test_model"

	s.AircraftRepository().Create(&models.Aircraft{
		Key:   1234,
		Model: model,
		Range: 234324,
	})

	aircraft, err := s.AircraftRepository().FindByMode(model)

	assert.NoError(t, err)
	assert.NotNil(t, aircraft)
}
