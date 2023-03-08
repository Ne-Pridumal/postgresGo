package store_test

import (
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBookingRepository(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("bookings")

	book, err := s.BookingRepository().Create(&models.Booking{
		BookRef:     1234,
		BookDate:    time.Now(),
		TotalAmount: 12312,
	})

	assert.NoError(t, err)
	assert.NotNil(t, book)
}

func TestBookingRepository_FindByRef(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("bookings")

	ref := 1234

	s.BookingRepository().Create(&models.Booking{
		BookRef:     int32(ref),
		BookDate:    time.Now(),
		TotalAmount: 123423,
	})

	book, err := s.BookingRepository().FindByRef(ref)

	assert.NoError(t, err)
	assert.NotNil(t, book)
}
