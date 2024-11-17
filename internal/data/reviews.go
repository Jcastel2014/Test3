package data

import (
	"context"
	"log"
	"time"
)

type ReviewIn struct {
	ID         int64     `json:"id"`
	Book_id    int64     `json:"book_id"`
	User_id    int64     `json:"user_id"`
	Review     string    `json:"review"`
	Created_at time.Time `json:"created_at"`
	Rating     float64   `json:"rating"`
}

func (b BookClub) InsertReview(review *ReviewIn) error {

	err := b.DoesBookExists(review.Book_id)

	if err != nil {
		return BookNotFound
	}

	err = b.DoesUserExists(review.User_id)

	if err != nil {
		return UserNotFound
	}

	// args := []any{book.Author}
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// err = b.DB.QueryRowContext(ctx, query, args...).Scan(&idA)

	query := `
	
	INSERT INTO book_reviews (book_id, user_id, review, rating, created_at) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id;
	
	`

	args := []any{review.Book_id, review.User_id, review.Review, review.Rating, review.Created_at}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = b.DB.QueryRowContext(ctx, query, args...).Scan(&review.ID)

	if err != nil {
		log.Println("error doing initla querry")
		return err
	}

	err = b.UpdateAverage(review.Book_id)

	if err != nil {
		log.Println("error doing update querry")
		return err
	}

	return nil

}
