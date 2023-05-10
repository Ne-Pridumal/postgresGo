package sqlstore

import (
	"database/sql"
)

// TODO: репоситорий для работы со всеми другими репами
type complexRepository struct {
	store *sql.DB
}

/*
func (r *complexRepository) Get(c *models.ComplexQuery) (models.Complex, error) {
	com := make(chan models.Complex)

	return <-com, nil
}

func (r *complexRepository) getAirportByCity(c chan models.Complex, city string) {
	query := `SELECT * FROM  airports, WHERE city = $1`
	var airportsSlice []models.Airport
	var a models.Airport
	rows, err := r.store.Query(
		query,
		city,
	)
	if err != nil {

	}
	for rows.Next() {
		if err := rows.Scan(
			&a.Key,
			&a.Name,
			&a.City,
			&a.Coordinates,
		); err != nil {
		}
		airportsSlice = append(airportsSlice, a)
	}

	c <-
}
*/
