package domain

import "net/http"

type Rester interface {
	healthcheckHandler(w http.ResponseWriter, r *http.Request)
	createMovieHandler(w http.ResponseWriter, r *http.Request)
	showMovieHandler(w http.ResponseWriter, r *http.Request)
	updateMovieHandler(w http.ResponseWriter, r *http.Request)
	deleteMovieHandler(w http.ResponseWriter, r *http.Request)
}
