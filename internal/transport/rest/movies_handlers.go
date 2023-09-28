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

	req := ProduceRequest{}

	err := app.Middleware.ReadJSON(w, r, &req)
	if err != nil {
		app.Status = http.StatusBadRequest
		app.Err = err
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &models.Movie{
		Title:   req.Title,
		Year:    req.Year,
		Runtime: req.Runtime,
		Genres:  req.Genres,
	}

	if models.ValidateMovie(app.Middleware.Validator, movie); !app.Middleware.Validator.Valid() {
		app.Status = http.StatusUnprocessableEntity
		app.Err = errors.New("validation failed")
		app.failedValidationResponse(w, r, app.Middleware.Validator.Errors)
		return
	}

	err = app.Storage.PGStore.Insert(movie)
	if err != nil {
		app.Status = http.StatusInternalServerError
		app.Err = err
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	err = app.Middleware.WriteJSON(w, http.StatusCreated, middleware.Envelope{"movie": movie}, headers)
	if err != nil {
		app.Status = http.StatusInternalServerError
		app.Err = err
		app.serverErrorResponse(w, r, err)
	}

	app.Status = http.StatusCreated
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
