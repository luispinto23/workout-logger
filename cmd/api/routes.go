package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	mux := mux.NewRouter()

	mux.NotFoundHandler = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowed)

	mux.Use(app.logAccess)
	mux.Use(app.recoverPanic)
	mux.Use(app.authenticate)

	mux.HandleFunc("/status", app.status).Methods("GET")

	mux.HandleFunc("/users", app.createUser).Methods("POST")
	mux.HandleFunc("/authentication-tokens", app.createAuthenticationToken).Methods("POST")

	authenticatedRoutes := mux.NewRoute().Subrouter()
	authenticatedRoutes.Use(app.requireAuthenticatedUser)
	authenticatedRoutes.HandleFunc("/protected", app.protected).Methods("GET")

	protectedRoutes := mux.NewRoute().Subrouter()
	protectedRoutes.Use(app.requireBasicAuthentication)
	protectedRoutes.HandleFunc("/basic-auth-protected", app.protected).Methods("GET")

	workoutRoutes := mux.NewRoute().Subrouter()
	workoutRoutes.HandleFunc("/workouts", app.createWorkoutHandler).Methods(http.MethodPost)
	workoutRoutes.HandleFunc("/workouts", app.listWorkoutsHandler).Methods(http.MethodGet)

	workoutRoutes.HandleFunc("/workouts/{id}/results", app.createResultHandler).Methods(http.MethodPost)

	exercisesRoutes := mux.NewRoute().Subrouter()
	exercisesRoutes.HandleFunc("/exercises", app.createExerciseHandler).Methods(http.MethodPost)
	exercisesRoutes.HandleFunc("/exercises/{id}/equipment", app.addExerciseEquipmentHandler).Methods(http.MethodPost)

	equipmentsRoutes := mux.NewRoute().Subrouter()
	equipmentsRoutes.HandleFunc("/equipments", app.createEquipmentHandler).Methods(http.MethodPost)

	return mux
}
