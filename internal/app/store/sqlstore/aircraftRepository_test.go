package sqlstore_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAircraftRepository(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("aircrafts")
	s := sqlstore.New(db)
	aircraft, err := s.Aircrafts().Create(&models.Aircraft{
		Key:   1234,
		Model: "123421342",
		Range: 234324,
	})

	assert.NoError(t, err)
	assert.NotNil(t, aircraft)
}

func TestAircraftRepository_FindByModel(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("aircrafts")
	s := sqlstore.New(db)
	model := "test_model"

	s.Aircrafts().Create(&models.Aircraft{
		Key:   1234,
		Model: model,
		Range: 234324,
	})

	aircraft, err := s.Aircrafts().FindByMode(model)

	assert.NoError(t, err)
	assert.NotNil(t, aircraft)
}
