CREATE TABLE ticket_flights(
  ticket_key text not null,
  flight_key bigserial not null,
  scheduled_departure timestamp not null,
  fare_condition text not null,
  amount bigserial not null,
  PRIMARY KEY(ticket_key, flight_key,scheduled_departure),
  CONSTRAINT fk_tickets FOREIGN KEY(ticket_key) REFERENCES tickets(ticket_key),
  CONSTRAINT fk_flights FOREIGN KEY(flight_key,scheduled_departure) REFERENCES flights(flight_key,scheduled_departure)
);

CREATE TABLE boarding_passes(
  ticket_key text not null,
  flight_key bigserial not null,
  scheduled_departure timestamp not null,
  boarding_no text not null,
  seat_no text not null,
  PRIMARY KEY(ticket_key, flight_key, scheduled_departure),
  CONSTRAINT fk_ticket_flights FOREIGN KEY(ticket_key, flight_key, scheduled_departure) REFERENCES ticket_flights(ticket_key, flight_key, scheduled_departure)
);
