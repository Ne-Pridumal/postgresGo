package store

import (
	"ne-pridumal/go-postgress/internal/app/models"
)

type AircraftRepository struct {
	store *Store
}

func (r *AircraftRepository) Create(a *models.Aircraft) (*models.Aircraft, error) {
	query := `
		INSERT INTO aircrafts (aircraft_key, model, range)
		VALUES ($1,$2,$3)
		RETURNING aircraft_key
	`
	if err := r.store.db.QueryRow(
		query,
		&a.Key, &a.Model, &a.Range,
	).Scan(&a.Key); err != nil {
		return nil, err
	}
	return a, nil
}
