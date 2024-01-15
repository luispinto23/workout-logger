CREATE TABLE IF NOT EXISTS exercises (
  id SERIAL NOT NULL PRIMARY KEY, 
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  type TEXT,
  muscle_group TEXT,
  difficulty TEXT,
  video_url TEXT
);