package domain

import "net/http"

type Rester interface {
	createMovieHandler(w http.ResponseWriter, r *http.Request)
	showMovieHandler(w http.ResponseWriter, r *http.Request)
}
