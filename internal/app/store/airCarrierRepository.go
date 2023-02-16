package store

import "ne-pridumal/go-postgress/internal/app/models"

type AirCarrierRepository struct {
	store *Store
}

func (r *AirCarrierRepository) Create(c *models.Air_carrier) (*models.Air_carrier, error) {
	query := "INSERT INTO air_carriers (air_carrier_key, name) VALUES ($1,$2) RETURNING air_carrier_key"
	if err := r.store.db.QueryRow(
		query,
		c.Air_carrier_key, c.Name,
	).Scan(&c.Air_carrier_key); err != nil {
		return nil, err
	}

	return c, nil
}

func (r *AirCarrierRepository) FindByName(name string) (*models.Air_carrier, error) {
	query := "SELECT * FROM air_carriers WHERE name=$1"

	c := &models.Air_carrier{}

	if err := r.store.db.QueryRow(
		query,
		name,
	).Scan(
		&c.Air_carrier_key,
		&c.Name,
	); err != nil {
		return nil, err
	}

	return c, nil
}
