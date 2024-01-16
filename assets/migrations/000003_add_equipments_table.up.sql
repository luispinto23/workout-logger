CREATE TABLE IF NOT EXISTS equipments (
    id SERIAL NOT NULL PRIMARY KEY,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL UNIQUE,
    description TEXT
);