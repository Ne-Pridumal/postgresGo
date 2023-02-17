package store

import (
	"ne-pridumal/go-postgress/internal/app/models"
)

type TicketRepository struct {
	store *Store
}

func (r *TicketRepository) Create(t *models.Ticket) (*models.Ticket, error) {
	query := "INSERT INTO tickets (flight_number, place, price, passenger_name, terminal, reg_time, takeoff_time, arrive_time) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING ticket_number"

	if err := r.store.db.QueryRow(
		query,
		t.FlightNumber, t.Place, t.Price, t.PassengerName, t.Terminal, t.RegTime, t.TakeoffTime, t.ArriveTime,
	).Scan(&t.TicketNumber); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *TicketRepository) FindByName(name string) (*models.Ticket, error) {
	query := "SELECT ticket_number, place, price, passenger_name, terminal, reg_time, takeoff_time, arrive_time FROM tickets WHERE passenger_name = $1"

	t := &models.Ticket{}

	if err := r.store.db.QueryRow(
		query,
		name,
	).Scan(
		&t.TicketNumber,
		&t.Place,
		&t.Price,
		&t.PassengerName,
		&t.Terminal,
		&t.RegTime,
		&t.TakeoffTime,
		&t.ArriveTime); err != nil {
		return nil, err
	}

	return t, nil
}
