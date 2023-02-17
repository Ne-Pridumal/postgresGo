package store

import "ne-pridumal/go-postgress/internal/app/models"

type SalesAgentRepository struct {
	store *Store
}

func (r *SalesAgentRepository) Create(a *models.SalesAgent) (*models.SalesAgent, error) {
	query := "INSERT INTO sales_agents (sales_agent_key, flight_key, sold_tickets_number) VALUES ($1, $2, $3) RETURNING sales_agent_key"
	if err := r.store.db.QueryRow(
		query,
		a.SalesAgentKey, a.FlightKey, a.SoldTicketsNumber,
	).Scan(&a.SalesAgentKey); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *SalesAgentRepository) FindByAgentKey(key uint) (*models.SalesAgent, error) {
	query := "SELECT * FROM sales_agents WHERE sales_agent_key=$1"

	a := &models.SalesAgent{}

	if err := r.store.db.QueryRow(
		query,
		key,
	).Scan(
		&a.SalesAgentKey,
		&a.FlightKey,
		&a.SoldTicketsNumber,
	); err != nil {
		return nil, err
	}

	return a, nil
}
