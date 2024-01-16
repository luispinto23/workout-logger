package database

import "context"

type Block struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Comment   string `db:"comment"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
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
