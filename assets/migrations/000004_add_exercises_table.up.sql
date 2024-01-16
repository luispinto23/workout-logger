CREATE TABLE IF NOT EXISTS exercises (
  id SERIAL NOT NULL PRIMARY KEY, 
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ,
  type TEXT,
  muscle_group TEXT,
  difficulty TEXT,
  video_url TEXT
);