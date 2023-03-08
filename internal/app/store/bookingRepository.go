package store

import (
	"ne-pridumal/go-postgress/internal/app/models"
)

type BookingRepository struct {
	store *Store
}

func (r *BookingRepository) Create(b *models.Booking) (*models.Booking, error) {
	query := `
	INSERT INTO bookings (book_ref, book_date, total_amount) 
	VALUES ($1,$2,$3) 
	RETURNING book_ref
	`

	if err := r.store.db.QueryRow(
		query,
		b.BookRef, b.BookDate, b.TotalAmount,
	).Scan(&b.BookRef); err != nil {
		return nil, err
	}
	return b, nil
}

func (r *BookingRepository) FindByRef(ref int) (*models.Booking, error) {
	query := `
		SELECT * FROM bookings
		WHERE book_ref = $1
	`

	b := &models.Booking{}

	if err := r.store.db.QueryRow(
		query,
		ref,
	).Scan(&b.BookRef, &b.BookDate, &b.TotalAmount); err != nil {
		return nil, err
	}

	return b, nil
}
