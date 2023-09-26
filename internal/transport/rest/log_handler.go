package rest

import (
	"errors"
	"net/http"
)

func (app *Application) logRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		app.Logger.Info().Msgf("[REQUEST] [%s] [%s] [%s] [%s] [%s]", r.UserAgent(), r.Proto, r.Method, r.Host, r.URL)

		next.ServeHTTP(w, r)

		if app.Status == 200 || app.Status == 201 {
			app.Logger.Info().Msgf("[RESPONSE] [%s] [%d]", r.Proto, app.Status)
			return
		}

		if app.Status == 404 || app.Status == 401 || app.Status == 422 {

			app.Logger.Error().Err(app.Err).Msgf("[RESPONSE] [%s] [%d]", r.Proto, app.Status)

			return
		}

		app.Status = 404

		app.Err = errors.New("not found")

		app.Logger.Error().Err(app.Err).Msgf("[RESPONSE] [%s] [%d]", r.Proto, app.Status)

	})
}
