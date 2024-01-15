package main

import (
	"my-workout-logs/internal/request"
	"my-workout-logs/internal/response"
	"net/http"
)

func (app *application) createResultHandler(w http.ResponseWriter, r *http.Request) {

	vars := request.GetVars(r)
	id := vars["id"]

	type partial struct {
		Weight int `json:"weight"`
		Reps   int `json:"reps"`
		Time   int `json:"time"`
		Series int `json:"series"`
	}

	var input struct {
		WorkoutID string    `json:"workout_id"`
		Partials  []partial `json:"partials"`
		Comment   string    `json:"comment"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	input.WorkoutID = id

	// send json response
	err = response.JSON(w, http.StatusCreated, input)
	if err != nil {
		app.serverError(w, r, err)
	}

}
