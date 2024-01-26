package main

import (
	"my-workout-logs/internal/database"
	"my-workout-logs/internal/request"
	"net/http"
	"strconv"
)

func (app *application) createExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		MuscleGroup string `json:"muscle_group"`
		Type        string `json:"type"`
		Difficulty  string `json:"difficulty"`
		VideoURL    string `json:"video_url"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	exercise := &database.Exercise{
		Name:         input.Name,
		Description:  input.Description,
		MuscleGroup:  input.MuscleGroup,
		Type:         input.Type,
		Difficulty:   input.Difficulty,
		DemoVideoURL: input.VideoURL,
	}

	_, err = app.db.InsertExercise(exercise)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) addExerciseEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := request.GetVars(r)
	exerciseID := vars["id"]

	intExerciseID, err := strconv.Atoi(exerciseID)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	var input struct {
		EquipmentID int `json:"equipment_id"`
	}

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	exerciseEquipment := &database.ExerciseEquipment{
		ExerciseID:  intExerciseID,
		EquipmentID: input.EquipmentID,
	}

	_, err = app.db.InsertExerciseEquipment(exerciseEquipment)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
