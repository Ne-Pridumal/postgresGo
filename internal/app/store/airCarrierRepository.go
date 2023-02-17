package store

import "ne-pridumal/go-postgress/internal/app/models"

type AirCarrierRepository struct {
	store *Store
}

func (r *AirCarrierRepository) Create(c *models.AirCarrier) (*models.AirCarrier, error) {
	query := "INSERT INTO air_carriers (air_carrier_key, name) VALUES ($1,$2) RETURNING air_carrier_key"
	if err := r.store.db.QueryRow(
		query,
		c.AirCarrierKey, c.Name,
	).Scan(&c.AirCarrierKey); err != nil {
		return nil, err
	}

	return c, nil
}

func (r *AirCarrierRepository) FindByName(name string) (*models.AirCarrier, error) {
	query := "SELECT * FROM air_carriers WHERE name= $1"

	c := &models.AirCarrier{}

	if err := r.store.db.QueryRow(
		query,
		name,
	).Scan(
		&c.AirCarrierKey,
		&c.Name,
	); err != nil {
		return nil, err
	}

	return c, nil
}
