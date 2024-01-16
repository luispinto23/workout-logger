CREATE TABLE workout_blocks (
  workout_id INT REFERENCES workouts(id),
  block_id INT REFERENCES blocks(id),
  PRIMARY KEY (workout_id, block_id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);