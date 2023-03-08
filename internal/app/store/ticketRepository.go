package store

import (
	"ne-pridumal/go-postgress/internal/app/models"
)

type TicketRepository struct {
	store *Store
}

func (r *TicketRepository) Create(t *models.Ticket) (*models.Ticket, error) {
	query := `
	INSERT INTO tickets (ticket_key, book_ref, passenger_id, passenger_name, contact_data) 
	VALUES ($1,$2,$3,$4,$5) 
	RETURNING ticket_key
	`

	if err := r.store.db.QueryRow(
		query,
		t.Key, t.BookRef, t.PassengerId, t.PassengerName, t.ContactDate,
	).Scan(&t.Key); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *TicketRepository) FindByName(name string) (*models.Ticket, error) {
	query := `
	SELECT * FROM tickets 
	WHERE passenger_name = $1
	`

	t := &models.Ticket{}

	if err := r.store.db.QueryRow(
		query,
		name,
	).Scan(
		&t.Key, &t.BookRef, &t.PassengerId, &t.PassengerName, &t.ContactDate,
	); err != nil {
		return nil, err
	}

	return t, nil
}
