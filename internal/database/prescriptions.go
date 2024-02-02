package database

import (
	"context"
	"database/sql"
	"time"
)

type Exercise struct {
	Name         string
	Description  string
	Type         string
	MuscleGroup  string
	Difficulty   int
	DemoVideoURL string
	Equipment    []string
}

type ExercisePrescription struct {
	Exercise Exercise
	MinReps  int
	MaxReps  int
	Weight   float64
	Duration int
	MinRest  int
	MaxRest  int
	Comment  string
}

type Block struct {
	Name        string
	Description string
	Exercises   []Exercise
}

type PrescriptionBlock struct {
	PrescriptionID int
	BlockID        int
	MinSeries      int
	MaxSeries      int
	Comment        string
}

type WorkoutPrescription struct {
	Name        string
	Description string
	Blocks      []PrescriptionBlock
}

type ExerciseResult struct {
	Exercise    Exercise
	Repetitions int
	Weight      float64
	Duration    int
	Comment     string
}

type BlockResult struct {
	Block           Block
	ExerciseResults []ExerciseResult
	Comment         string
}

type WorkoutResult struct {
	WorkoutPrescriptionID int
	BlockResults          []BlockResult
	Comment               string
}

func (db *DB) InsertWorkoutPrescription(tx *sql.Tx, wp *WorkoutPrescription) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int

	// Insert prescription into the database
	query := `
        INSERT INTO workout_prescriptions (name, description, created_at)
        VALUES ($1, $2, $3)
        RETURNING id`

	// Execute the query and scan the returned id into the id variable
	err := tx.QueryRowContext(ctx, query, wp.Name, wp.Description, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	// Insert each block associated with the prescription into the database
	for _, block := range wp.Blocks {
		// Insert prescription_block record
		err := db.InsertPrescriptionBlock(tx, id, block.BlockID, block.MinSeries, block.MaxSeries, block.Comment)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (db *DB) InsertPrescriptionBlock(tx *sql.Tx, prescriptionID, blockID, minSeries, maxSeries int, comment string) error {
	// Insert prescription_block record into the database
	query := `
        INSERT INTO prescription_blocks (prescription_id, block_id, min_series, max_series, comment)
        VALUES ($1, $2, $3, $4, $5)`

	_, err := tx.Exec(query, prescriptionID, blockID, minSeries, maxSeries, comment)
	return err
}
