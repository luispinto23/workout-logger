package database

import (
	"context"
	"my-workout-logs/internal/request"
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

func (db *DB) CreateWorkoutPrescription(prescription *request.WorkoutPrescriptionRequest) (int, error) {
	// Create a new workout prescription
	workoutPrescription := WorkoutPrescription{
		Name:        prescription.Name,
		Description: prescription.Description,
		Blocks:      []WorkoutBlock{},
		created_at:  time.Now(),
		updated_at:  time.Now(),
	}

	// Create workout blocks
	for _, block := range prescription.Blocks {
		blockID, err := db.CreateWorkoutBlock(&WorkoutBlock{
			Name:        block.Name,
			Description: block.Description,
			Exercises:   []Exercise{},
			MinSeries:   block.MinSeries,
			MaxSeries:   block.MaxSeries,
			Comment:     block.Comment,
			created_at:  time.Now(),
			updated_at:  time.Now(),
		})
		if err != nil {
			return 0, err
		}
		workoutPrescription.Blocks = append(workoutPrescription.Blocks, WorkoutBlock{ID: blockID})
	}

	// Insert workout prescription into the database
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	var id int
	query := `
    INSERT INTO workout_prescriptions (name, description, blocks, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id`

	err := db.GetContext(ctx, &id, query, workoutPrescription.Name, workoutPrescription.Description, workoutPrescription.Blocks, time.Now())
	if err != nil {
		return 0, err
	}

	// Update workout prescription with its ID
	workoutPrescription.ID = id

	return id, nil
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
