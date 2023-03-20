package sqlstore

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/models"
	"strconv"
)

type bookingRepository struct {
	store *sql.DB
}

func (r *bookingRepository) Create(b *models.Booking) (*models.Booking, error) {
	query := `
	INSERT INTO bookings (book_ref, book_date, total_amount) 
	VALUES ($1,$2,$3) 
	RETURNING book_ref
	`

	if err := r.store.QueryRow(
		query,
		b.BookRef, b.BookDate, b.TotalAmount,
	).Scan(&b.BookRef); err != nil {
		return nil, err
	}
	return b, nil
}

func (r *bookingRepository) GetAll() ([]models.Booking, error) {
	query := "SELECT * FROM bookings"
	rows, err := r.store.Query(query)
	var bookingSlice []models.Booking
	var b models.Booking

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&b.BookRef,
			&b.BookDate,
			&b.TotalAmount,
		); err != nil {
			return nil, err
		}
		bookingSlice = append(bookingSlice, b)
	}

	return bookingSlice, nil
}

func (r *bookingRepository) FindByRef(ref int) (*models.Booking, error) {
	query := `
		SELECT * FROM bookings
		WHERE book_ref = $1
	`

	b := &models.Booking{}

	if err := r.store.QueryRow(
		query,
		ref,
	).Scan(&b.BookRef, &b.BookDate, &b.TotalAmount); err != nil {
		return nil, err
	}

	return b, nil
}

func (r *bookingRepository) Delete(id int) error {
	fields := map[string]string{
		"book_ref": strconv.FormatInt(int64(id), 10),
	}
	return deleteByFields(fields, "bookings", r.store)
}
