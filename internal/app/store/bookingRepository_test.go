package store_test

import (
	"fmt"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBookingRepository(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("bookings")
	s := store.New(db)
	book, err := s.BookingRepository().Create(&models.Booking{
		BookRef:     1234,
		BookDate:    time.Now(),
		TotalAmount: 12312,
	})

	assert.NoError(t, err)
	assert.NotNil(t, book)
}

func TestBookingRepository_GetAll(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("bookings")
	s := store.New(db)

	ref := 1234

	s.BookingRepository().Create(&models.Booking{
		BookRef:     ref,
		BookDate:    time.Now(),
		TotalAmount: 123423,
	})
	s.BookingRepository().Create(&models.Booking{
		BookRef:     343243,
		BookDate:    time.Now(),
		TotalAmount: 123423,
	})
	bookings, err := s.BookingRepository().GetAll()

	assert.NotNil(t, bookings)
	assert.NoError(t, err)
	fmt.Print(bookings)
}

func TestBookingRepository_FindByRef(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("bookings")
	s := store.New(db)
	ref := 1234

	s.BookingRepository().Create(&models.Booking{
		BookRef:     ref,
		BookDate:    time.Now(),
		TotalAmount: 123423,
	})

	book, err := s.BookingRepository().FindByRef(ref)

	assert.NoError(t, err)
	assert.NotNil(t, book)
}
