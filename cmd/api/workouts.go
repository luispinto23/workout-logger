package main

import (
	"my-workout-logs/internal/request"
	"net/http"
)

func (app *application) createWorkoutPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	var newPrescription request.WorkoutPrescriptionRequest

	err := request.DecodeJSON(w, r, &newPrescription)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	_, err = app.db.createWorkoutPrescription(newPrescription)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) listWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
}
