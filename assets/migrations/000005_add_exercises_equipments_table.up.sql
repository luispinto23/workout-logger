CREATE TABLE IF NOT EXISTS exercises_equipments (
    exercise_id INTEGER NOT NULL REFERENCES exercises(id),
    equipment_id INTEGER NOT NULL REFERENCES equipments(id),
    PRIMARY KEY (exercise_id, equipment_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);