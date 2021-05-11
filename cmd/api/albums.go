package main

import (
	"encoding/json"
	"fmt"
	"github.com/tklara86/greenlight/internal/data"
	"net/http"
	"time"
)

// createMovieHandler POST request to /v1/movies endpoint
func (app *application) createAlbumHandler(w http.ResponseWriter, r *http.Request) {
	
	var input struct{
		Title 	 	string `json:"title"`
		Year  	  	int32 `json:"year"`
		Runtime 	int32 `json:"runtime"`
		Artist		[]string `json:"artist"`
		Genres		[]string `json:"genres"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		app.errorResponse(w,r, http.StatusBadRequest, err.Error())
	}

	_, err = fmt.Fprintf(w, "%+v\n", input)
	if err != nil {
		app.errorResponse(w,r, http.StatusBadRequest, "Error creating an album")
	}
}

// showMovieHandler GET request to /v1/movies/:id endpoint
func (app *application) showAlbumHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	movie := data.Album{
		ID: id,
		Title: "Actions",
		CreatedAt: time.Now(),
		Runtime: 102,
		Artist: []string{"Krzysztof Penderecki", "Don Cherry"},
		Genres: []string{"avant-garde", "classical", "jazz"},
		Version: 1,
	}

	err = app.writeToJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}