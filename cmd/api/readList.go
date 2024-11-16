package main

import (
	"fmt"
	"net/http"

	"github.com/Jcastel2014/test3/internal/data"
	"github.com/Jcastel2014/test3/internal/validator"
)

func (a *appDependencies) postReadingList(w http.ResponseWriter, r *http.Request) {

	var incomingData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Created_by  string `json:"created_by"`
	}

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	readList := &data.ReadList{
		Name:        incomingData.Name,
		Description: incomingData.Description,
		Created_by:  incomingData.Created_by,
	}

	v := validator.New()
	data.ValidateList(v, readList)

	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.bookclub.InsertList(readList)

	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/list/%d", readList.ID))

	data := envelope{
		"readList": readList,
	}

	err = a.writeJSON(w, http.StatusCreated, data, headers)

	if err != nil {
		a.serverErrResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", incomingData)

}
