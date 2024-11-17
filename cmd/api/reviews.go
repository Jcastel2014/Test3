package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Jcastel2014/test3/internal/data"
	"github.com/Jcastel2014/test3/internal/validator"
)

func (a *appDependencies) postReview(w http.ResponseWriter, r *http.Request) {

	id, err := a.readIDParam(r)

	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	_, err = a.bookclub.GetBook(id)

	if err != nil {
		a.bookNotFound(w, r, err)
	}

	var incomingData struct {
		User_id    int64   `json:"user_id"`
		Review     string  `json:"review"`
		Created_at string  `json:"created_at"`
		Rating     float64 `json:"rating"`
	}

	err = a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	review := &data.ReviewIn{
		Book_id:    id,
		User_id:    incomingData.User_id,
		Review:     incomingData.Review,
		Created_at: time.Now(),
		Rating:     incomingData.Rating,
	}

	v := validator.New()
	data.ValidateReview(v, review)

	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.bookclub.InsertReview(review)

	if err != nil {
		a.hello()
		a.serverErrResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/review/%d", review.ID))

	data := envelope{
		"review": review,
	}

	err = a.writeJSON(w, http.StatusCreated, data, headers)

	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", incomingData)

}
