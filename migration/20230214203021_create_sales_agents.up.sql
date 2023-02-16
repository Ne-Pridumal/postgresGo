CREATE TABLE sales_agents(
  sales_agent_key bigserial not null primary key,
  flight_key bigserial not null,
  sold_tickets_number bigserial not null
);
