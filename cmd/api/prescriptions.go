package main

import (
	"my-workout-logs/internal/database"
	"my-workout-logs/internal/request"
	"my-workout-logs/internal/response"
	"net/http"
)

func (app *application) getPrescriptions(w http.ResponseWriter, r *http.Request) {
}

func (app *application) getPrescription(w http.ResponseWriter, r *http.Request) {
}

func (app *application) createPrescription(w http.ResponseWriter, r *http.Request) {
	var input database.WorkoutPrescription

	// Decode JSON request body into WorkoutPrescription struct
	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	// Begin transaction
	tx, err := app.db.Begin()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			app.serverError(w, r, err)
		}
	}()

	// Insert prescription into the database
	prescriptionID, err := app.db.InsertWorkoutPrescription(tx, &input)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Send JSON response with status code 201 (Created)
	responseBody := map[string]interface{}{
		"id": prescriptionID,
	}
	err = response.JSON(w, http.StatusCreated, responseBody)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updatePrescription(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deletePrescription(w http.ResponseWriter, r *http.Request) {
}

func (app *application) getPrescriptionResults(w http.ResponseWriter, r *http.Request) {
}

func (app *application) createPrescriptionResult(w http.ResponseWriter, r *http.Request) {
}
