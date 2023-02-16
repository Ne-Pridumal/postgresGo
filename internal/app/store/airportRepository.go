package store

import (
	"ne-pridumal/go-postgress/internal/app/models"
)

type AirportRepository struct {
	store *Store
}

func (r *AirportRepository) Create(a *models.Airport) (*models.Airport, error) {
	query := "INSERT INTO airports (address, name) VALUES ($1, $2) RETURNING airport_key"

	if err := r.store.db.QueryRow(
		query,
		a.Address, a.Name,
	).Scan(&a.Airport_key); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AirportRepository) FindByName(name string) (*models.Airport, error) {
	query := "SELECT airport_key, name, address FROM airports WHERE name = $1"

	a := &models.Airport{}

	if err := r.store.db.QueryRow(
		query,
		name,
	).Scan(
		&a.Airport_key,
		&a.Address,
		&a.Name,
	); err != nil {
		return nil, err
	}

	return a, nil
}
