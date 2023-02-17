package store

import "ne-pridumal/go-postgress/internal/app/models"

type AirportAgentRepository struct {
	store *Store
}

func (r *AirportAgentRepository) Create(a *models.AirportAgent) (*models.AirportAgent, error) {
	query := "INSERT INTO airport_agents (airport_agent_key, name, airport_key, air_carrier_key) VALUES ($1,$2,$3,$4) RETURNING airport_agent_key"
	if err := r.store.db.QueryRow(
		query,
		a.AirCarrierKey, a.Name, a.AirportKey, a.AirportAgentKey,
	).Scan(
		&a.AirportAgentKey,
	); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *AirportAgentRepository) FindByName(name string) (*models.AirportAgent, error) {
	query := "SELECT * FROM airport_agents WHERE name = $1"

	a := &models.AirportAgent{}
	if err := r.store.db.QueryRow(
		query,
		name,
	).Scan(
		&a.AirportAgentKey,
		&a.Name,
		&a.AirportKey,
		&a.AirCarrierKey,
	); err != nil {
		return nil, err
	}

	return a, nil
}
