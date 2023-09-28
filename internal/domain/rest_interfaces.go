package domain

import "net/http"

type Rester interface {
	healthcheckHandler(w http.ResponseWriter, r *http.Request)
	createMovieHandler(w http.ResponseWriter, r *http.Request)
	showMovieHandler(w http.ResponseWriter, r *http.Request)
}
