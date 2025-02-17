package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.nikolasavic.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//
	// movie := data.Movie{
	// 	Title:   input.Title,
	// 	Year:    input.Year,
	// 	Runtime: data.Runtime(input.Runtime),
	// 	Genres:  input.Genres,
	// }
	//
	// err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// }

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Matrix",
		Year:      1999,
		Runtime:   136,
		Genres:    []string{"action", "scifi"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
