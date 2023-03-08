CREATE TABLE flights(
  flight_key bigserial not null,
  scheduled_departure timestamp not null,
  scheduled_arrival timestamp not null,
  departure_airport bigserial not null,
  arrival_airport bigserial not null,
  status text not null,
  aircraft_key bigserial not null,
  actual_departure timestamp not null,
  actual_arrival timestamp not null,
  PRIMARY KEY(flight_key, scheduled_departure)
);
