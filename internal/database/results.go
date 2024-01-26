package database

import (
	"context"
	"time"
)

type blockResult struct {
	BlockID         int    `json:"block_id"`
	WorkoutResultID int    `json:"workout_result_id"`
	Comment         string `json:"comment"`
}

type WorkoutResultRequest struct {
	Date            string           `json:"date"`
	ExerciseResults []ExerciseResult `json:"exercise_results"`
	BlockResults    []BlockResult    `json:"block_results"`
	Comment         string           `json:"comment"`
}

func (db *DB) InsertResult(newResult *WorkoutResultRequest, workoutID int) (int, error) {
	// // Validate exercise results
	// for _, exerciseResult := range newResult.ExerciseResults {
	// 	if err := app.validateExerciseResult(exerciseResult); err != nil {
	// 		return err
	// 	}
	// }
	//
	// // Validate block results
	// for _, blockResult := range newResult.BlockResults {
	// 	if err := app.validateBlockResult(blockResult); err != nil {
	// 		return err
	// 	}
	// }

	// Create workout result
	workoutResult := WorkoutResult{
		Date:            newResult.Date,
		ExerciseResults: newResult.ExerciseResults,
		BlockResults:    newResult.BlockResults,
		Comment:         newResult.Comment,
		WorkoutID:       workoutID,
	}

	// Insert workout result into the database
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int
	query := `
    INSERT INTO workout_results (date, exercise_results, block_results, comment, workout_id, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id`

	err := db.GetContext(ctx, &id, query, workoutResult.Date, workoutResult.ExerciseResults, workoutResult.BlockResults, workoutResult.Comment, workoutID, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	// Update workout result with its ID
	workoutResult.ID = id

	return id, err
}
