package store

import "ne-pridumal/go-postgress/internal/app/models"

type SalesAgentRepository struct {
	store *Store
}

func (r *SalesAgentRepository) Create(a *models.Sales_agent) (*models.Sales_agent, error) {
	query := "INSERT INTO sales_agents (sales_agent_key, flight_key, sold_tickets_number) VALUES ($1, $2, $3) RETURNING sales_agent_key"
	if err := r.store.db.QueryRow(
		query,
		a.Sales_agent_key, a.Flight_key, a.Sold_tickets_number,
	).Scan(&a.Sales_agent_key); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *SalesAgentRepository) FindByAgentKey(key uint) (*models.Sales_agent, error) {
	query := "SELECT * FROM sales_agents WHERE sales_agent_key=$1"

	a := &models.Sales_agent{}

	if err := r.store.db.QueryRow(
		query,
		key,
	).Scan(
		&a.Sales_agent_key,
		&a.Flight_key,
		&a.Sold_tickets_number,
	); err != nil {
		return nil, err
	}

	return a, nil
}
