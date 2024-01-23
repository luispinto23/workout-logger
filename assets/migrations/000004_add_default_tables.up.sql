CREATE TABLE exercises (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  type VARCHAR(255) NOT NULL,
  muscle_group VARCHAR(255) NOT NULL,
  difficulty VARCHAR(255) NOT NULL,
  demo_video_url TEXT,
  equipment_list TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE workout_blocks (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  min_series INTEGER,
  max_series INTEGER,
  comment TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE workout_prescriptions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE exercises_in_workout_prescriptions (
  exercise_id INTEGER NOT NULL REFERENCES exercises(id),
  workout_prescription_id INTEGER NOT NULL REFERENCES workout_prescriptions(id),
  min_reps INTEGER,
  max_reps INTEGER,
  weight NUMERIC(10,2),
  duration NUMERIC(10,2),
  min_rest INTEGER,
  max_rest INTEGER,
  comment TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (exercise_id, workout_prescription_id)
);

CREATE TABLE blocks_in_workout_prescriptions (
  block_id INTEGER NOT NULL REFERENCES workout_blocks(id),
  workout_prescription_id INTEGER NOT NULL REFERENCES workout_prescriptions(id),
  min_series INTEGER,
  max_series INTEGER,
  comment TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (block_id, workout_prescription_id)
);

CREATE TABLE workout_results (
  id SERIAL PRIMARY KEY,
  workout_prescription_id INTEGER NOT NULL REFERENCES workout_prescriptions(id),
  date DATE NOT NULL,
  comment TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE exercise_results_in_workout_results (
  exercise_id INTEGER NOT NULL REFERENCES exercises(id),
  workout_result_id INTEGER NOT NULL REFERENCES workout_results(id),
  repetitions INTEGER,
  weight NUMERIC(10,2),
  duration NUMERIC(10,2),
  comment TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (exercise_id, workout_result_id)
);

CREATE TABLE block_results_in_workout_results (
  block_id INTEGER NOT NULL REFERENCES workout_blocks(id),
  workout_result_id INTEGER NOT NULL REFERENCES workout_results(id),
  comment TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (block_id, workout_result_id)
);
