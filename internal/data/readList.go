package data

import (
	"context"
	"fmt"
	"log"
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

func (b BookClub) GetAllLists(filters Filters) ([]*ReadList, error) {
	query := fmt.Sprintf(`
	SELECT R.id, R.name, R.description, U.user_name AS created_by, S.name as status 
	FROM readList AS R 
	INNER JOIN users AS U 
	ON R.created_by = U.id 
	INNER JOIN status AS S 
	ON R.status = S.id
	ORDER BY %s %s, R.id ASC
	LIMIT $1 OFFSET $2
	`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := b.DB.QueryContext(ctx, query, filters.limit(), filters.offset())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	readLists := []*ReadList{}

	for rows.Next() {
		var readList ReadList
		err := rows.Scan(&readList.ID, &readList.Name, &readList.Description, &readList.Created_by, &readList.Status)
		if err != nil {
			return nil, err
		}

		readLists = append(readLists, &readList)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return readLists, nil
}

func (b BookClub) ListAddBook(id int64, bid int64) error {
	err := b.DoesBookExists(bid)

	if err != nil {
		log.Println("Book not found with id %d", bid)
		return err
	}

	err = b.DoesListExists(id)

	if err != nil {
		fmt.Sprintf("lsit not found %d", id)
		return err
	}

	query := `
	INSERT INTO book_list (book_id, list_id)
	VALUES ($1, $2)
	RETURNING id
	
	`

	args := []any{bid, id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return b.DB.QueryRowContext(ctx, query, args...).Scan(&id)

}
