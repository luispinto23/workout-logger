package main

import (
	"my-workout-logs/internal/database"
	"my-workout-logs/internal/request"
	"net/http"
)

func (app *application) createEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	equipment := &database.Equipment{
		Name: input.Name,
	}

	_, err = app.db.InsertEquipment(equipment)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
