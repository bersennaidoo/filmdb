package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bersennaidoo/filmdb/internal/domain/models"
	"github.com/bersennaidoo/lib/pkg/middleware"
)

func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	app.Status = http.StatusCreated
	fmt.Fprintln(w, "create a new movie")
}

func (app *Application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.Status = http.StatusNotFound
		app.Err = err
		http.NotFound(w, r)
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
		app.Err = err
		http.Error(w, "The server encountered a problem and could not process your request",
			http.StatusInternalServerError)
	}

	app.Status = http.StatusOK
}
