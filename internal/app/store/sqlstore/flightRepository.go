package sqlstore

import (
	"database/sql"
	"ne-pridumal/go-postgress/internal/app/models"
)

type FlightRepository struct {
	store *sql.DB
}

func (r *FlightRepository) Create(f *models.Flight) (*models.Flight, error) {
	query := `
		INSERT INTO flights 
			(flight_key, 
			scheduled_departure, 
			scheduled_arrival, 
			departure_airport, 
			arrival_airport, 
			status, 
			aircraft_key, 
			actual_departure, 
			actual_arrival) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) 
		RETURNING flight_key
	`

	if err := r.store.QueryRow(
		query,
		&f.FlightKey,
		&f.ScheduledDeparture,
		&f.ScheduledArrival,
		&f.DepartureAirport,
		&f.ArrivalAirport,
		&f.Status,
		&f.AircraftKey,
		&f.ActualDeparture,
		&f.ActualArrival,
	).Scan(&f.FlightKey); err != nil {
		return nil, err
	}
	return f, nil
}

func (r *FlightRepository) GetAll() ([]models.Flight, error) {
	query := "SELECT * from flights"
	rows, err := r.store.Query(query)
	var flightsSlice []models.Flight
	var f models.Flight

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&f.FlightKey,
			&f.ScheduledDeparture,
			&f.ScheduledArrival,
			&f.DepartureAirport,
			&f.ArrivalAirport,
			&f.Status,
			&f.AircraftKey,
			&f.ActualDeparture,
			&f.ActualArrival,
		); err != nil {
			return nil, err
		}
		flightsSlice = append(flightsSlice, f)
	}

	return flightsSlice, nil
}

func (r *FlightRepository) FindByFlightKey(key int) (*models.Flight, error) {
	query := `
		SELECT * FROM flights 
		WHERE flight_key = $1
	`

	f := &models.Flight{}

	if err := r.store.QueryRow(
		query,
		key,
	).Scan(
		&f.FlightKey,
		&f.ScheduledDeparture,
		&f.ScheduledArrival,
		&f.DepartureAirport,
		&f.ArrivalAirport,
		&f.Status,
		&f.AircraftKey,
		&f.ActualDeparture,
		&f.ActualArrival,
	); err != nil {
		return nil, err
	}
	return f, nil
}

func (r *FlightRepository) DeleteById(id int) error {
	return deleteByIdTableWithIdField(id, "flights", "flight_key", r.store)
}
