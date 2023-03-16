package sqlstore

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/models"
)

type AirportRepository struct {
	store *sql.DB
}

func (r *AirportRepository) Create(a *models.Airport) (*models.Airport, error) {
	query := `
		INSERT INTO airports (airport_key, airport_name, city, coordinates) 
		VALUES ($1,$2,$3,$4)
		RETURNING airport_key
	`

	if err := r.store.QueryRow(
		query,
		&a.Key, &a.Name, &a.City, &a.Coordinates,
	).Scan(&a.Key); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AirportRepository) GetAll() ([]models.Airport, error) {
	query := "SELECT * FROM airports"
	rows, err := r.store.Query(query)
	var airportsSlice []models.Airport
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var a models.Airport

		if err := rows.Scan(
			&a.Key,
			&a.Name,
			&a.City,
			&a.Coordinates,
		); err != nil {
			return nil, err
		}
		airportsSlice = append(airportsSlice, a)
	}

	return airportsSlice, err
}

func (r *AirportRepository) FindByName(name string) (*models.Airport, error) {
	query := `
		SELECT airport_key, airport_name, city, coordinates FROM airports 
		WHERE airport_name = $1
	`

	a := &models.Airport{}

	if err := r.store.QueryRow(
		query,
		name,
	).Scan(
		&a.Key,
		&a.Name,
		&a.City,
		&a.Coordinates,
	); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *AirportRepository) DeleteById(id int) error {
	return deleteByIdTableWithIdField(id, "airports", "airport_key", r.store)
}
