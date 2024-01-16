CREATE TABLE block_exercises (
  block_id INT REFERENCES blocks(id),
  exercise_properties_id INT REFERENCES block_exercise_properties(id),
  PRIMARY KEY (block_id, exercise_properties_id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);