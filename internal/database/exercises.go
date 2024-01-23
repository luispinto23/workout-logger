package database

import (
	"context"
	"time"
)

// Exercise struct
type Exercise struct {
	ID            int    `db:"id"`
	Name          string `db:"name"`
	Description   string `db:"description"`
	Type          string `db:"type"`
	MuscleGroup   string `db:"muscle_group"`
	Difficulty    string `db:"difficulty"`
	DemoVideoURL  string `db:"demo_video_url"`
	EquipmentList string `db:"equipment_list"`
}

// ExerciseInWorkoutPrescription struct
type ExerciseInWorkoutPrescription struct {
	ExerciseID int     `db:"exercise_id"`
	WorkoutID  int     `db:"workout_id"`
	MinReps    int     `db:"min_reps"`
	MaxReps    int     `db:"max_reps"`
	Weight     float32 `db:"weight"`
	Duration   float32 `db:"duration"`
	MinRest    int     `db:"min_rest"`
	MaxRest    int     `db:"max_rest"`
	Comment    string  `db:"comment"`
}

// ExerciseResult struct
type ExerciseResult struct {
	ExerciseID      int     `db:"exercise_id"`
	WorkoutResultID int     `db:"workout_result_id"`
	Repetitions     int     `db:"repetitions"`
	Weight          float32 `db:"weight"`
	Duration        float32 `db:"duration"`
	Comment         string  `db:"comment"`
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
