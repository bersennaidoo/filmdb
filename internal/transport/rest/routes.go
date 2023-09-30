package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.HealthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.CreateMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.ShowMovieHandler)
	router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.UpdateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.DeleteMovieHandler)

	return app.logRequest(router)
}
