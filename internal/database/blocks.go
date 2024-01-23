package database

import "context"

// BlockInWorkoutPrescription struct
type BlockInWorkoutPrescription struct {
	BlockID   int    `db:"block_id"`
	WorkoutID int    `db:"workout_id"`
	MinSeries int    `db:"min_series"`
	MaxSeries int    `db:"max_series"`
	Comment   string `db:"comment"`
}

// BlockResult struct
type BlockResult struct {
	BlockID         int    `db:"block_id"`
	WorkoutResultID int    `db:"workout_result_id"`
	Comment         string `db:"comment"`
}

func (db *DB) InsertBlock(block *Block) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int

	query := `
		INSERT INTO blocks (name, comment)
		VALUES ($1, $2)
		RETURNING id`

	err := db.GetContext(ctx, &id, query, block.Name, block.Comment)
	if err != nil {
		return 0, err
	}

	return id, err
}
