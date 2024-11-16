package data

import (
	"context"
	"log"
	"time"

	"github.com/Jcastel2014/test3/internal/validator"
)

func (b BookClub) DoesAuthorExists(author string) (error, int) {
	query := `
		SELECT id
		FROM authors
		WHERE name = $1
	`
	args := []any{author}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int

	err := b.DB.QueryRowContext(ctx, query, args...).Scan(&id)

	log.Println(id)

	if err != nil {
		return err, -1
	}

	return nil, id
}

func ValidateBook(v *validator.Validator, book *Book) {

	v.Check(book.Title != "", "title", "must be provided")
	v.Check(len(book.Title) <= 255, "title", "must not be more than 100 byte long")

	v.Check(book.ISBN != "", "isbn", "must be provided")
	v.Check(len(book.ISBN) <= 255, "isbn", "must not be more than 100 byte long")

	v.Check(book.Author != "", "author", "must be provided")
	v.Check(len(book.Author) <= 100, "author", "must not be more than 100 characters long")

	v.Check(book.Genre != "", "genre", "must be provided")
	v.Check(len(book.Genre) <= 50, "genre", "must not be more than 50 characters long")

	v.Check(len(book.Description) <= 1000, "description", "must not be more than 1000 characters long")

	v.Check(!book.Publication_Date.IsZero(), "publication_date", "must be provided")
	v.Check(book.Publication_Date.Before(time.Now()), "publication_date", "must not be in the future")

	// v.Check(review.Rating > 0, "rating", "must be greater than 0")
	// v.Check(review.Rating <= 5, "rating", "must be less than 5")
	// v.Check(len(review.Comment) <= 100, "comment", "must not be more than 100 byte long")

}

func ValidateList(v *validator.Validator, list *ReadList) {

	v.Check(list.Name != "", "name", "must be provided")
	v.Check(len(list.Name) <= 255, "name", "must not be more than 100 byte long")

	v.Check(len(list.Description) <= 1000, "description", "must not be more than 1000 characters long")

	v.Check(list.Created_by != "", "created_by", "must be provided")
	v.Check(len(list.Created_by) <= 100, "created_by", "must not be more than 100 characters long")

	// v.Check(list.Status != "", "status", "must be provided")
	// v.Check(len(list.Status) <= 50, "status", "must not be more than 50 characters long")

	// v.Check(list.Status == "Completed" || list.Status == "Currently Reading", "status", "must be Completed or Currently Reading")

}
