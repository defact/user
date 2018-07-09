CREATE TABLE profile_connections (
  id SERIAL PRIMARY KEY,
  profile_id INTEGER,
  connected_to INTEGER,
  type VARCHAR(32)
);