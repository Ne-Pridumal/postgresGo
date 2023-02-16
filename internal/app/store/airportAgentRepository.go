package store

import "ne-pridumal/go-postgress/internal/app/models"

type AirportAgentRepository struct {
	store *Store
}

func (r *AirportAgentRepository) Create(a *models.Airport_agent) (*models.Airport_agent, error) {
	query := "INSERT INTO airport_agents (airport_agent_key, name, airport_key, air_carrier_key) VALUES ($1,$2,$3,$4) RETURNING airport_agent_key"
	if err := r.store.db.QueryRow(
		query,
		a.Air_carrier_key, a.Name, a.Airport_key, a.Airport_agent_key,
	).Scan(
		&a.Airport_agent_key,
	); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *AirportAgentRepository) FindByName(name string) (*models.Airport_agent, error) {
	query := "SELECT * FROM airport_agents WHERE name = $1"

	a := &models.Airport_agent{}
	if err := r.store.db.QueryRow(
		query,
		name,
	).Scan(
		&a.Airport_agent_key,
		&a.Name,
		&a.Airport_key,
		&a.Air_carrier_key,
	); err != nil {
		return nil, err
	}

	return a, nil
}
