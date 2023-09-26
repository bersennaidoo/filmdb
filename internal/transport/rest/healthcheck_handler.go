package rest

import (
	"net/http"

	"github.com/bersennaidoo/lib/pkg/middleware"
)

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := middleware.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.Config.Environment,
			"version":     version,
		},
	}

	err := app.Middleware.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.Status = http.StatusInternalServerError
		app.Err = err
		app.serverErrorResponse(w, r, err)
	}

	app.Status = http.StatusOK
}
