package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bersennaidoo/filmdb/internal/domain/models"
	"github.com/bersennaidoo/filmdb/internal/service/storage"
	"github.com/bersennaidoo/lib/pkg/middleware"
)

func (app *Application) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {

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

func (app *Application) ShowMovieHandler(w http.ResponseWriter, r *http.Request) {

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

func (app *Application) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {

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

	reqBody := PatchRequest{}

	err = app.Middleware.ReadJSON(w, r, &reqBody)
	if err != nil {
		app.Status = http.StatusBadRequest
		app.Err = err
		app.badRequestResponse(w, r, err)
		return
	}

	if reqBody.Title != nil {
		movie.Title = *reqBody.Title
	}

	if reqBody.Year != nil {
		movie.Year = *reqBody.Year
	}

	if reqBody.Runtime != nil {
		movie.Runtime = *reqBody.Runtime
	}

	if reqBody.Genres != nil {
		movie.Genres = reqBody.Genres
	}

	if models.ValidateMovie(app.Middleware.Validator, movie); !app.Middleware.Validator.Valid() {
		app.Status = http.StatusBadRequest
		app.Err = errors.New("bad request")
		app.failedValidationResponse(w, r, app.Middleware.Validator.Errors)
		return
	}

	err = app.Storage.PGStore.Update(movie)
	if err != nil {
		switch {
		case error.Is(err, storage.ErrEditConflict):
			app.Status = http.StatusConflict
			app.Err = err
			app.editConflictResponse(w, r)
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

func (app *Application) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {

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
