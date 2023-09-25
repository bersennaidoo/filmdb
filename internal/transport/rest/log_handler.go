package rest

import "net/http"

func (app *Application) logRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		app.Logger.Info().Msgf("[REQUEST] [%s] [%s] [%s] [%s] [%s]", r.UserAgent(), r.Proto, r.Method, r.Host, r.URL)

		next.ServeHTTP(w, r)

		app.Logger.Info().Err(app.Err).Msgf("[RESPONSE] [%s] [%d]", r.Proto, app.Status)

	})
}
