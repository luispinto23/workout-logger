CREATE TABLE blocks (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  comment TEXT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);