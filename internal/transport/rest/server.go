package rest

import (
	"net/http"
	"time"
)

func (app *Application) ServerRun() {

	srv := &http.Server{
		Addr:         app.Config.Server.Conn,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Info().Msgf("starting %s server on %s", app.Config.Environment,
		app.Config.Server.Conn)

	app.Logger.Info().Msg("[LISTEN GET /v1/healthcheck]")
	app.Logger.Info().Msg("[LISTEN POST /v1/movies]")
	app.Logger.Info().Msg("[LISTEN GET /v1/movies/:id]")

	err := srv.ListenAndServe()

	app.Logger.Info().Err(err).Msg("")
}
