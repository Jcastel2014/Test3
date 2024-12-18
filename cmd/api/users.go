package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/Jcastel2014/test3/internal/data"
	"github.com/Jcastel2014/test3/internal/validator"
)

func (a *appDependencies) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	var incomingData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
	}

	user := &data.User{
		Username:  incomingData.Username,
		Email:     incomingData.Email,
		Activated: false,
	}

	//hashing password and storing with the cleartect version
	err = user.Password.Set(incomingData.Password)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	//user validation
	v := validator.New()
	data.ValidateUser(v, user)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.userModel.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			a.failedValidationResponse(w, r, v.Errors)
		default:
			a.serverErrResponse(w, r, err)
		}
		return
	}

	//new activation token to expire in 3 days time
	token, err := a.tokenModel.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	data := envelope{
		"user": user,
	}

	/*	Send the email as a Goroutine. We do this because it might take a long time
		and we don't want our handler to wait for that to finish. We will implement
		the background() function later
	*/
	a.background(func() {
		data := map[string]any{
			"activationToken": token.PlainText,
			"userID":          user.ID,
		}
		err = a.mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			a.logger.Error(err.Error())
		}

	})

	//status code 201 resource created
	err = a.writeJSON(w, http.StatusCreated, data, nil)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}
}

func (a *appDependencies) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	//get the body from the request and store in temporary struct

	var incomingData struct {
		TokenPlainText string `json:"token"`
	}

	err := a.readJSON(w, r, &incomingData)

	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	//validate the data
	v := validator.New()
	data.ValidatetokenPlaintext(v, incomingData.TokenPlainText)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	//verification of token belonging to the user
	user, err := a.userModel.GetForToken(data.ScopeActivation, incomingData.TokenPlainText)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("token", "invalid or expired activation token")
			a.failedValidationResponse(w, r, v.Errors)
		default:
			a.serverErrResponse(w, r, err)
		}
		return
	}

	//correct token provided so we activate them
	user.Activated = true
	err = a.userModel.Update(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			a.editConflictResponse(w, r)
		default:
			a.serverErrResponse(w, r, err)
		}
		return
	}

	//delete actiavtion token after user activation
	err = a.tokenModel.DeleteAllForUser(data.ScopeActivation, user.ID)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	//send response
	data := envelope{
		"user": user,
	}
	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}
}

// func (a *appDependencies) listUserProfileHandler(w http.ResponseWriter, r *http.Request) {
// 	//get the id from the URL so that we can use it to query the comments table.
// 	//'uid' for userID
// 	id, err := a.readIDParam(r, "uid")
// 	if err != nil {
// 		a.notFoundResponse(w, r)
// 		return
// 	}

// 	//call the GetUserProfile() function to retrieve
// 	user, err := a.userModel.GetByID(id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			a.notFoundResponse(w, r)
// 		default:
// 			a.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}

// 	//display the user information
// 	data := envelope{
// 		"user": user,
// 	}

// 	err = a.writeJSON(w, http.StatusOK, data, nil)
// 	if err != nil {
// 		a.serverErrorResponse(w, r, err)
// 		return
// 	}
// }
