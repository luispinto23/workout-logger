package main

import (
	"my-workout-logs/internal/request"
	"my-workout-logs/internal/response"
	"net/http"
)

func (app *application) createWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	type Equipment struct {
		Name string `json:"name"`
	}

	type exercise struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		MuscleGroup string `json:"muscle_group"`
		Type        string `json:"type"`
		Equipment   []Equipment
		Difficulty  string `json:"difficulty"`
		SeriesMin   int    `json:"series_min"`
		SeriesMax   int    `json:"series_max"`
		RepMin      int    `json:"rep_min"`
		RepMax      int    `json:"rep_max"`
		RestMin     int    `json:"rest_min"`
		RestMax     int    `json:"rest_max"`
		Weight      int    `json:"weight"`
		Duration    int    `json:"duration"`
		VideoURL    string `json:"video_url"`
		Comment     string `json:"comment"`
	}

	type block struct {
		Name      string     `json:"name"`
		Exercises []exercise `json:"exercises"`
		Comment   string     `json:"comment"`
	}

	var input struct {
		Name    string  `json:"name"`
		Date    string  `json:"date"`
		Blocks  []block `json:"blocks"`
		Comment string  `json:"comment"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	// send json response
	err = response.JSON(w, http.StatusCreated, input)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) listWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
}
