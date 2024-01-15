CREATE TABLE block_exercise_properties (
  id SERIAL PRIMARY KEY,
  exercise_id INT REFERENCES exercises(id),
  series_min INT,
  series_max INT,
  rep_min INT,
  rep_max INT,
  rest_min INT,
  rest_max INT,
  weight INT,
  duration INT,
  comment TEXT
);