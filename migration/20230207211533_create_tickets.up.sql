CREATE TABLE tickets (
  ticket_number bigserial not null primary key,
  flight_number bigserial not null,
  place text not null,
  price bigserial not null,
  passenger_name text not null,
  terminal serial not null,
  reg_time timestamp not null,
  takeoff_time timestamp not null,
  arrive_time timestamp not null
);
