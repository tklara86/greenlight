package main

import (
	"fmt"
	"github.com/tklara86/greenlight/internal/data"
	"net/http"
	"time"
)

// createMovieHandler POST request to /v1/movies endpoint
func (app *application) createAlbumHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Create a new album")
	if err != nil {
		app.notFoundResponse(w, r)
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