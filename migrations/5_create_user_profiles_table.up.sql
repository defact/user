CREATE TABLE user_profiles (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  profile_id INTEGER,
  is_primary BOOLEAN
);