package data

import (
	"context"
	"time"
)

type ReadList struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Created_by  string `json:"created_by"`
	Status      string `json:"status"`
}

func (b BookClub) InsertList(readList *ReadList) error {

	query := `
	INSERT INTO readList(name, description, created_by, status)
	VALUES ($1, $2, $3, 1)
	RETURNING id
	
	`

	args := []any{readList.Name, readList.Description, 1}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return b.DB.QueryRowContext(ctx, query, args...).Scan(&readList.ID)

}
