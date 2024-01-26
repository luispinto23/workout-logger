package main

import (
	"my-workout-logs/internal/database"
	"my-workout-logs/internal/request"
	"net/http"
	"strconv"
)

func (app *application) createResultHandler(w http.ResponseWriter, r *http.Request) {
	vars := request.GetVars(r)
	id := vars["id"]

	workoutId, err := strconv.Atoi(id)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	var resultInput database.WorkoutResultRequest
	err = request.DecodeJSON(w, r, &resultInput)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	_, err = app.db.InsertResult(&resultInput, workoutId)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
