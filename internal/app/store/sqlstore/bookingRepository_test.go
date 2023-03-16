package sqlstore_test

import (
	"fmt"
	"ne-pridumal/go-postgress/internal/app/models"
	"ne-pridumal/go-postgress/internal/app/store/sqlstore"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBookingRepository(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("bookings")
	s := sqlstore.New(db)
	book, err := s.Bookings().Create(&models.Booking{
		BookRef:     1234,
		BookDate:    time.Now(),
		TotalAmount: 12312,
	})

	assert.NoError(t, err)
	assert.NotNil(t, book)
}

func TestBookingRepository_GetAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("bookings")
	s := sqlstore.New(db)

	ref := 1234

	s.Bookings().Create(&models.Booking{
		BookRef:     ref,
		BookDate:    time.Now(),
		TotalAmount: 123423,
	})
	s.Bookings().Create(&models.Booking{
		BookRef:     343243,
		BookDate:    time.Now(),
		TotalAmount: 123423,
	})
	bookings, err := s.Bookings().GetAll()

	assert.NotNil(t, bookings)
	assert.NoError(t, err)
	fmt.Print(bookings)
}

func TestBookingRepository_FindByRef(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("bookings")
	s := sqlstore.New(db)
	ref := 1234

	s.Bookings().Create(&models.Booking{
		BookRef:     ref,
		BookDate:    time.Now(),
		TotalAmount: 123423,
	})

	book, err := s.Bookings().FindByRef(ref)

	assert.NoError(t, err)
	assert.NotNil(t, book)
}
