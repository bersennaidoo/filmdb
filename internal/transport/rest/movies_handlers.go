package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bersennaidoo/filmdb/internal/domain/models"
	"github.com/bersennaidoo/filmdb/internal/service/storage"
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

	movie, err := app.Storage.PGStore.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrRecordNotFound):
			app.Status = http.StatusNotFound
			app.Err = err
			app.notFoundResponse(w, r)
		default:
			app.Status = http.StatusInternalServerError
			app.Err = err
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.Middleware.WriteJSON(w, http.StatusOK, middleware.Envelope{"movie": movie}, nil)
	if err != nil {
		app.Status = http.StatusInternalServerError
		app.Err = err
		app.serverErrorResponse(w, r, err)
	}

	app.Status = http.StatusOK
}

func (app *Application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.Status = http.StatusNotFound
		app.Err = err
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.Storage.PGStore.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrRecordNotFound):
			app.Status = http.StatusNotFound
			app.Err = err
			app.notFoundResponse(w, r)
		default:
			app.Status = http.StatusInternalServerError
			app.Err = err
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	reqBody := ProduceRequest{}

	err = app.Middleware.ReadJSON(w, r, &reqBody)
	if err != nil {
		app.Status = http.StatusBadRequest
		app.Err = err
		app.badRequestResponse(w, r, err)
		return
	}

	movie.Title = reqBody.Title
	movie.Year = reqBody.Year
	movie.Runtime = reqBody.Runtime
	movie.Genres = reqBody.Genres

	if models.ValidateMovie(app.Middleware.Validator, movie); !app.Middleware.Validator.Valid() {
		app.Status = http.StatusBadRequest
		app.Err = errors.New("bad request")
		app.failedValidationResponse(w, r, app.Middleware.Validator.Errors)
		return
	}

	err = app.Storage.PGStore.Update(movie)
	if err != nil {
		app.Status = http.StatusInternalServerError
		app.Err = err
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.Middleware.WriteJSON(w, http.StatusOK, middleware.Envelope{"movie": movie}, nil)
	if err != nil {
		app.Status = http.StatusInternalServerError
		app.Err = err
		app.serverErrorResponse(w, r, err)
	}

	app.Status = http.StatusOK
}

func (app *Application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.Status = http.StatusNotFound
		app.Err = err
		app.notFoundResponse(w, r)
		return
	}

	err = app.Storage.PGStore.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrRecordNotFound):
			app.Status = http.StatusNotFound
			app.Err = err
			app.notFoundResponse(w, r)
		default:
			app.Status = http.StatusInternalServerError
			app.Err = err
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.Middleware.WriteJSON(w, http.StatusOK, middleware.Envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.Status = http.StatusInternalServerError
		app.Err = err
		app.serverErrorResponse(w, r, err)
	}

	app.Status = http.StatusOK
}
