package store

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/models"
)

type AircraftRepository struct {
	store *sql.DB
}

func (r *AircraftRepository) Create(a *models.Aircraft) (*models.Aircraft, error) {
	query := `
		INSERT INTO aircrafts (aircraft_key, model, range)
		VALUES ($1,$2,$3)
		RETURNING aircraft_key
	`
	if err := r.store.QueryRow(
		query,
		&a.Key, &a.Model, &a.Range,
	).Scan(&a.Key); err != nil {
		return nil, err
	}
	return a, nil
}

func (r *AircraftRepository) GetAll() ([]models.Aircraft, error) {
	query := "SELECT * FROM aircrafts"
	rows, err := r.store.Query(query)
	var aircraftsSlice []models.Aircraft
	var a models.Aircraft
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&a.Key,
			&a.Model,
			&a.Range,
		); err != nil {
			return nil, err
		}
		aircraftsSlice = append(aircraftsSlice, a)
	}

	return aircraftsSlice, nil
}

func (r *AircraftRepository) FindByMode(modelName string) (*models.Aircraft, error) {
	query := `
		SELECT aircraft_key, model, range FROM aircrafts
		WHERE model = $1
	`
	a := &models.Aircraft{}

	if err := r.store.QueryRow(
		query,
		modelName,
	).Scan(&a.Key, &a.Model, &a.Range); err != nil {
		return nil, err
	}

	return a, nil
}
