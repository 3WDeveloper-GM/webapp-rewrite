package main

// This file stores the routes that are made when moving on the website.

import (
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// I need this function to get the routes of the http.ServeMux object.
// It was funny, if you didn't use the configuration generated from the database
// You cannot do the rest of the work.
func Router(app *config.Application) http.Handler {

	r := mux.NewRouter()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileserver)).Methods(http.MethodGet)

	r.HandleFunc("/", Home(app)).Methods(http.MethodGet)
	r.HandleFunc("/snippet/view/{id:[0-9]+}", View(app)).Methods(http.MethodGet)
	r.HandleFunc("/snippet/create", Wcreate(app)).Methods(http.MethodGet, http.MethodPost)

	//Chaining some middleware
	standard := alice.New(app.RecoverPanic, app.LogRequest, config.SecureHeaders)

	return standard.Then(r)
}
