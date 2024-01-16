package database

import (
	"context"
	"time"
)

type Workout struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Name      string    `db:"name"`
	Date      time.Time `db:"date"`
	Comment   string    `db:"comment"`
	UpdatedAt time.Time `db:"updated_at"`
	UserID    int       `db:"user_id"`
}

type WorkoutBlock struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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
