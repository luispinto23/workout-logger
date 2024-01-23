package database

import (
	"context"
	"time"
)

// WorkoutBlock struct
type WorkoutBlock struct {
	ID          int        `db:"id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	MinSeries   int        `db:"min_series"`
	MaxSeries   int        `db:"max_series"`
	Comment     string     `db:"comment"`
	Exercises   []Exercise `db:"exercises"`
}

// WorkoutPrescription struct
type WorkoutPrescription struct {
	ID          int            `db:"id"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	Blocks      []WorkoutBlock `db:"blocks"`
}

// WorkoutResult struct
type WorkoutResult struct {
	ID              int              `db:"id"`
	WorkoutID       int              `db:"workout_id"`
	Date            string           `db:"date"`
	ExerciseResults []ExerciseResult `db:"exercise_results"`
	BlockResults    []BlockResult    `db:"block_results"`
	Comment         string           `db:"comment"`
}

func (db *DB) InsertWorkout(workout *Workout) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int
	now := time.Now()

	query := `
		INSERT INTO workouts (name, date, comment, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err := db.GetContext(ctx, &id, query, workout.Name, workout.Date, workout.Comment, workout.UserID, now)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (db *DB) InsertWorkoutBlock(block *Block) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int

	query := `
		INSERT INTO blocks (created_at, name, comment)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := db.GetContext(ctx, &id, query, time.Now(), block.Name, block.Comment)
	if err != nil {
		return 0, err
	}

	return id, err
}
