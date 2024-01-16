package database

import (
	"context"
	"time"
)

type Equipment struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func (db *DB) InsertEquipment(equipment *Equipment) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	var id int
	now := time.Now()

	query := `
		INSERT INTO equipment (name, description, created_at, updated_at)
		VALUES ($1, $2)
		RETURNING id`

	err := db.GetContext(ctx, &id, query, equipment.Name, equipment.Description, now, now)
	if err != nil {
		return 0, err
	}

	return id, err
}
