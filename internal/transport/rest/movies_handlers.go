package rest

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bersennaidoo/filmdb/internal/domain/models"
	"github.com/bersennaidoo/lib/pkg/middleware"
)

func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string             `json:"title"`
		Year    int32              `json:"year"`
		Runtime middleware.Runtime `json:"runtime"`
		Genres  []string           `json:"genres"`
	}

	err := app.Middleware.ReadJSON(w, r, &input)
	if err != nil {
		app.Status = http.StatusBadRequest
		app.Err = err
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &models.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	if models.ValidateMovie(app.Middleware.Validator, movie); !app.Middleware.Validator.Valid() {
		app.Status = http.StatusUnprocessableEntity
		app.Err = errors.New("validation failed")
		app.failedValidationResponse(w, r, app.Middleware.Validator.Errors)
		return
	}

	app.Status = http.StatusCreated
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *Application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.Status = http.StatusNotFound
		app.Err = err
		app.notFoundResponse(w, r)
		return
	}

	movie := models.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.Middleware.WriteJSON(w, http.StatusOK, middleware.Envelope{"movie": movie}, nil)
	if err != nil {
		app.Status = http.StatusInternalServerError
		app.Err = err
		app.serverErrorResponse(w, r, err)
	}

	app.Status = http.StatusOK
}
