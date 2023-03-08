CREATE TABLE seats(
  aircraft_key bigserial not null,
  seat_no text not null,
  fare_conditions text not null,
  PRIMARY KEY(aircraft_key, seat_no),
  CONSTRAINT fk_aircrafts FOREIGN KEY(aircraft_key) REFERENCES aircrafts(aircraft_key)
);
