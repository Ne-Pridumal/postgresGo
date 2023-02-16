CREATE TABLE airplanes(
  airplane_key bigserial not null primary key,
  number_of_sits serial not null,
  model text not null,
  air_carrier_key bigserial not null
);
