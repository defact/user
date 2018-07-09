CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(128) UNIQUE NOT NULL,
  hashed_password VARCHAR(32),
  salt VARCHAR(32),
  verification_code VARCHAR(16),
  networks JSON,
  is_locked BOOLEAN,
  is_archived BOOLEAN,
  settings JSON
);