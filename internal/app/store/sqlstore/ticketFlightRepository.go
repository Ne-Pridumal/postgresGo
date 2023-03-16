package sqlstore

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/models"
)

type TicketFlightRepository struct {
	store *sql.DB
}

func (r *TicketFlightRepository) Create(t *models.TicketFlights) (*models.TicketFlights, error) {
	query := `
		INSERT INTO ticket_flights
		(
			ticket_key,
			flight_key,
			scheduled_departure,
			fare_condition,
			amount
		)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING ticket_key
	`

	if err := r.store.QueryRow(
		query,
		&t.TicketKey,
		&t.FlightKey,
		&t.ScheduledDeparture,
		&t.FareCondition,
		&t.Amount,
	).Scan(&t.TicketKey); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *TicketFlightRepository) GetAll() ([]models.TicketFlights, error) {
	query := "SELECT * FROM ticket_flights"
	rows, err := r.store.Query(query)
	var ticketFlightsSlice []models.TicketFlights
	var t models.TicketFlights

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&t.TicketKey,
			&t.FlightKey,
			&t.ScheduledDeparture,
			&t.FareCondition,
			&t.Amount,
		); err != nil {
			return nil, err
		}

		ticketFlightsSlice = append(ticketFlightsSlice, t)
	}

	return ticketFlightsSlice, nil
}

func (r *TicketFlightRepository) DeleteById(id int) error {
	return deleteByIdTableWithIdField(id, "ticket_flights", "ticket_key", r.store)
}
