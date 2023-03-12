package store

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/models"
)

type BoardingPassRepository struct {
	store *sql.DB
}

func (r *BoardingPassRepository) Create(bp *models.BoardingPass) (*models.BoardingPass, error) {
	query := `
		INSERT INTO boarding_passes
		(
			ticket_key,
			flight_key,
			scheduled_departure,
			boarding_no,
			seat_no
		)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING ticket_key, flight_key
	`

	if err := r.store.QueryRow(
		query,
		&bp.TicketKey,
		&bp.FlightKey,
		&bp.ScheduledDeparture,
		&bp.BoardingNo,
		&bp.SeatNo,
	).Scan(&bp.TicketKey, &bp.FlightKey); err != nil {
		return nil, err
	}

	return bp, nil
}

func (r *BoardingPassRepository) GetAll() ([]models.BoardingPass, error) {
	query := "SELECT * FROM boarding_passes"
	rows, err := r.store.Query(query)
	var boardingPassesSlice []models.BoardingPass
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bp models.BoardingPass

		if err := rows.Scan(
			&bp.TicketKey,
			&bp.FlightKey,
			&bp.ScheduledDeparture,
			&bp.BoardingNo,
			&bp.SeatNo,
		); err != nil {
			return nil, err
		}

		boardingPassesSlice = append(boardingPassesSlice, bp)
	}

	return boardingPassesSlice, nil
}
