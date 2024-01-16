package database

import (
	"context"
	"time"
)

type Exercise struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	MuscleGroup string `db:"muscle_group"`
	Type        string `db:"type"`
	Difficulty  string `db:"difficulty"`
	VideoURL    string `db:"video_url"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

type ExerciseEquipment struct {
	ID          int `db:"id"`
	ExerciseID  int `db:"exercise_id"`
	EquipmentID int `db:"equipment_id"`
	CreatedAt   int `db:"created_at"`
	UpdatedAt   int `db:"updated_at"`
}

func (db *DB) InsertExercise(exercise *Exercise) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int
	now := time.Now()

	query := `
		INSERT INTO exercises (name, description, muscle_group, type, difficulty, video_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := db.GetContext(ctx, &id, query, exercise.Name, exercise.Description, exercise.MuscleGroup, exercise.Type, exercise.Difficulty, exercise.VideoURL, now, now)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (db *DB) InsertExerciseEquipment(ee *ExerciseEquipment) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int
	now := time.Now()

	query := `
		INSERT INTO exercises_equipments (exercise_id, equipment_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err := db.GetContext(ctx, &id, query, ee.ExerciseID, ee.EquipmentID, now, now)

	if err != nil {
		return 0, err
	}

	return id, err
}
