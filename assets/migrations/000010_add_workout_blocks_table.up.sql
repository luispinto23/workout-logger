CREATE TABLE workout_blocks (
  workout_id INT REFERENCES workouts(id),
  block_id INT REFERENCES blocks(id),
  PRIMARY KEY (workout_id, block_id),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);