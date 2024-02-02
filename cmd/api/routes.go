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

	prescriptionRoutes := mux.NewRoute().Subrouter()
	prescriptionRoutes.HandleFunc("/prescriptions", app.getPrescriptions).Methods("GET")
	prescriptionRoutes.HandleFunc("/prescriptions/{id}", app.getPrescription).Methods("GET")
	prescriptionRoutes.HandleFunc("/prescriptions", app.createPrescription).Methods("POST")
	prescriptionRoutes.HandleFunc("/prescriptions/{id}", app.updatePrescription).Methods("PUT")
	prescriptionRoutes.HandleFunc("/prescriptions/{id}", app.deletePrescription).Methods("DELETE")

	prescriptionRoutes.HandleFunc("/prescriptions/{id}/results", app.getPrescriptionResults).Methods("GET")
	prescriptionRoutes.HandleFunc("/prescriptions/{id}/results", app.createPrescriptionResult).Methods("POST")
	return mux
}
