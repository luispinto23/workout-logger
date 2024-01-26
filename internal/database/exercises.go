package database

import (
	"context"
	"time"
)

// Exercise struct
type Exercise struct {
	ID            int    `db:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Type          string `json:"type"`
	MuscleGroup   string `json:"muscle_group"`
	Difficulty    string `json:"difficulty"`
	DemoVideoURL  string `json:"demo_video_url"`
	EquipmentList string `json:"equipment_list"`
}

// ExerciseInWorkoutPrescription struct
type ExerciseInWorkoutPrescription struct {
	ExerciseID int     `json:"exercise_id"`
	WorkoutID  int     `json:"workout_id"`
	MinReps    int     `json:"min_reps"`
	MaxReps    int     `json:"max_reps"`
	Weight     float32 `json:"weight"`
	Duration   float32 `json:"duration"`
	MinRest    int     `json:"min_rest"`
	MaxRest    int     `json:"max_rest"`
	Comment    string  `json:"comment"`
}

// ExerciseResult struct
type ExerciseResult struct {
	ExerciseID      int     `json:"exercise_id"`
	WorkoutResultID int     `json:"workout_result_id"`
	Repetitions     int     `json:"repetitions"`
	Weight          float32 `json:"weight"`
	Duration        float32 `json:"duration"`
	Comment         string  `json:"comment"`
}

type ExerciseEquipment struct {
	ID          int `db:"id"`
	ExerciseID  int `json:"exercise_id"`
	EquipmentID int `json:"equipment_id"`
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

	err := db.GetContext(ctx, &id, query, exercise.Name, exercise.Description, exercise.MuscleGroup, exercise.Type, exercise.Difficulty, exercise.DemoVideoURL, now, now)
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
