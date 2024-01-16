CREATE TABLE workouts (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  date DATE,
  comment TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  user_id INT REFERENCES users(id)
);