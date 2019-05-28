CREATE TABLE IF NOT EXISTS auth (
  id SERIAL PRIMARY KEY,
  user_uuid UUID NOT NULL UNIQUE,
  password text NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);