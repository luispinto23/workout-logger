package database

import "context"

// BlockInWorkoutPrescription struct
type BlockInWorkoutPrescription struct {
	BlockID   int    `json:"block_id"`
	WorkoutID int    `json:"workout_id"`
	MinSeries int    `json:"min_series"`
	MaxSeries int    `json:"max_series"`
	Comment   string `json:"comment"`
}

// BlockResult struct
type BlockResult struct {
	BlockID         int    `json:"block_id"`
	WorkoutResultID int    `json:"workout_result_id"`
	Comment         string `json:"comment"`
}

func (db *DB) InsertBlock(block *BlockInWorkoutPrescription) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int

	query := `
		INSERT INTO blocks (name, comment)
		VALUES ($1, $2)
		RETURNING id`

	err := db.GetContext(ctx, &id, query, block.BlockID, block.WorkoutID)
	if err != nil {
		return 0, err
	}

	return id, err
}
