package sqlstore

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/models"
	"strconv"
)

type seatRepository struct {
	store *sql.DB
}

func (r *seatRepository) Create(s *models.Seat) (*models.Seat, error) {
	query := `
		INSERT INTO seats (aircraft_key, seat_no, fare_conditions)
		VALUES ($1,$2,$3)
		RETURNING aircraft_key
	`

	if err := r.store.QueryRow(
		query,
		&s.AircraftKey, &s.Number, &s.FareConditions,
	).Scan(&s.AircraftKey); err != nil {
		return nil, err
	}

	return s, nil
}

func (r *seatRepository) GetAll() ([]models.Seat, error) {
	query := "SELECT * FROM seats"
	rows, err := r.store.Query(query)
	var seatsSlice []models.Seat
	var s models.Seat

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&s.AircraftKey,
			&s.Number,
			&s.FareConditions,
		); err != nil {
			return nil, err
		}

		seatsSlice = append(seatsSlice, s)
	}

	return seatsSlice, nil
}

func (r *seatRepository) Delete(aCode int, no string) error {
	fields := map[string]string{
		"aircraft_key": strconv.FormatInt(int64(aCode), 10),
		"seat_no":      no,
	}
	return deleteByFields(fields, "seats", r.store)
}
