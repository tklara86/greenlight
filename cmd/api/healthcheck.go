package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := envelope{
		"status": "Available",
		"system_info": map[string]string{
			"Environment": app.config.env,
			"Version": version,
		},
	}

	err := app.writeToJson(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
