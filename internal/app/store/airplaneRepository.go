package store

import (
	"ne-pridumal/go-postgress/internal/app/models"
)

type AirplaneRepository struct {
	store *Store
}

func (r *AirplaneRepository) Create(a *models.Airplane) (*models.Airplane, error) {
	query := "INSERT INTO airplanes (airplane_key, number_of_sits, model, air_carrier_key) VALUES ($1,$2,$3,$4) RETURNING airplane_key"
	if err := r.store.db.QueryRow(
		query,
		a.Airplane_key, a.Number_of_sits, a.Model, a.Air_carrier_key,
	).Scan(&a.Airplane_key); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *AirplaneRepository) FindByModel(model string) (*models.Airplane, error) {
	query := "SELECT airplane_key, number_of_sits, model, air_carrier_key FROM airplanes WHERE model = $1"

	a := &models.Airplane{}

	if err := r.store.db.QueryRow(
		query,
		model,
	).Scan(
		&a.Airplane_key,
		&a.Number_of_sits,
		&a.Model,
		&a.Air_carrier_key,
	); err != nil {
		return nil, err
	}

	return a, nil
}
