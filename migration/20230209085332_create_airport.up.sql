CREATE TABLE airports(
  airport_key bigserial not null primary key,
  airport_name text not null,
  city text not null,
  coordinates text not null
);
