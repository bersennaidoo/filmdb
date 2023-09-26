package rest

import (
	"fmt"
	"net/http"

	"github.com/bersennaidoo/lib/pkg/middleware"
)

func (app *Application) logError(r *http.Request, err error) {

	app.Logger.Error().Err(err).Msg("")
}

func (app *Application) errorResponse(w http.ResponseWriter, r *http.Request, status int,
	message interface{}) {

	env := middleware.Envelope{"error": message}

	err := app.Middleware.WriteJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *Application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *Application) notFoundResponse(w http.ResponseWriter, r *http.Request) {

	message := "the requested resource could not be found"

	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *Application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {

	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)

	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
