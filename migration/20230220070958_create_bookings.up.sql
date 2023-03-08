CREATE TABLE bookings(
  book_ref bigserial not null primary key,
  book_date timestamp not null,
  total_amount bigserial not null
);

CREATE TABLE tickets(
  ticket_key text not null primary key,
  book_ref bigserial not null,
  passenger_id bigserial not null,
  passenger_name text not null,
  contact_data timestamp not null,
  CONSTRAINT fk_booking FOREIGN KEY(book_ref) REFERENCES bookings(book_ref)
);
