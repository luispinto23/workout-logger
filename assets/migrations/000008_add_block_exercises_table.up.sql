CREATE TABLE block_exercises (
  block_id INT REFERENCES blocks(id),
  exercise_properties_id INT REFERENCES block_exercise_properties(id),
  PRIMARY KEY (block_id, exercise_properties_id),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);