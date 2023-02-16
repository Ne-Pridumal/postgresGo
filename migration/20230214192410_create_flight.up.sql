CREATE TABLE flights(
  flight_key bigserial not null primary key,
  airplane_key bigserial not null,
  starting_airport_key bigserial not null,
  final_airport_key bigserial not null,
  max_tickets_number bigserial not null,
  sales_agent_key bigserial not null,
  air_carrier_key bigserial not null,
  takeoff_time timestamp not null,
  arrive_time timestamp not null,
  date timestamp not null
);
