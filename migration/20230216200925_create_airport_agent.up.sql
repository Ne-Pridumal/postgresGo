CREATE TABLE airport_agents(
  airport_agent_key bigserial not null primary key,
  name text not null,
  airport_key bigserial not null,
  air_carrier_key bigserial not null 
);
