CREATE TABLE IF NOT EXISTS auth (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL,
  password text NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);