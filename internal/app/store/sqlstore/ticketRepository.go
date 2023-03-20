package sqlstore

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/models"
	"strconv"
)

type ticketRepository struct {
	store *sql.DB
}

func (r *ticketRepository) Create(t *models.Ticket) (*models.Ticket, error) {
	query := `
	INSERT INTO tickets (ticket_key, book_ref, passenger_id, passenger_name, contact_data) 
	VALUES ($1,$2,$3,$4,$5) 
	RETURNING ticket_key
	`

	if err := r.store.QueryRow(
		query,
		&t.Key,
		&t.BookRef,
		&t.PassengerId,
		&t.PassengerName,
		&t.ContactDate,
	).Scan(&t.Key); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *ticketRepository) GetAll() ([]models.Ticket, error) {
	query := "SELECT * FROM tickets"
	rows, err := r.store.Query(query)
	var ticketsSlice []models.Ticket
	var t models.Ticket

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		if err := rows.Scan(
			&t.Key,
			&t.BookRef,
			&t.PassengerId,
			&t.PassengerName,
			&t.ContactDate,
		); err != nil {
			return nil, err
		}

		ticketsSlice = append(ticketsSlice, t)
	}
	return ticketsSlice, nil
}

func (r *ticketRepository) FindByName(name string) (*models.Ticket, error) {
	query := `
	SELECT * FROM tickets 
	WHERE passenger_name = $1
	`

	t := &models.Ticket{}

	if err := r.store.QueryRow(
		query,
		name,
	).Scan(
		&t.Key, &t.BookRef, &t.PassengerId, &t.PassengerName, &t.ContactDate,
	); err != nil {
		return nil, err
	}

	return t, nil
}

func (r *ticketRepository) Delete(key string, ref int) error {
	fields := map[string]string{
		"ticket_key": key,
		"book_ref":   strconv.FormatInt(int64(ref), 10),
	}
	return deleteByFields(fields, "tickets", r.store)
}
