package store

import "ne-pridumal/go-postgress/internal/app/models"

type FlightRepository struct {
	store *Store
}

func (r *FlightRepository) Create(f *models.Flight) (*models.Flight, error) {
	query := "INSERT INTO flights (flight_key, airplane_key, starting_airport_key, final_airport_key, max_tickets_number, sales_agent_key, air_carrier_key, takeoff_time, arrive_time, date) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9, $10) RETURNING flight_key"

	if err := r.store.db.QueryRow(
		query,
		f.Flight_key, f.Airplane_key, f.Starting_airport_key, f.Final_airport_key, f.Max_tickets_number, f.Sales_agent_key, f.Air_carrier_key, f.Takeoff_time, f.Arrive_time, f.Date,
	).Scan(&f.Flight_key); err != nil {
		return nil, err
	}
	return f, nil
}

func (r *FlightRepository) FindByFlightKey(key uint) (*models.Flight, error) {
	query := "SELECT flight_key, airplane_key, starting_airport_key, final_airport_key, max_tickets_number, sales_agent_key, air_carrier_key, takeoff_time, arrive_time, date from flights WHERE flight_key = $1"

	f := &models.Flight{}

	if err := r.store.db.QueryRow(
		query,
		key,
	).Scan(
		&f.Flight_key,
		&f.Airplane_key,
		&f.Starting_airport_key,
		&f.Final_airport_key,
		&f.Max_tickets_number,
		&f.Sales_agent_key,
		&f.Air_carrier_key,
		&f.Takeoff_time,
		&f.Arrive_time,
		&f.Date,
	); err != nil {
		return nil, err
	}
	return f, nil
}
