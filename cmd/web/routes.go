package main

// This file stores the routes that are made when moving on the website.

import (
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg"
	"github.com/gorilla/mux"
)

// I need this function to get the routes of the http.ServeMux object.
// It was funny, if you didn't use the configuration generated from the database
// You cannot do the rest of the work.
func Router(app *config.Application) *mux.Router {

	r := mux.NewRouter()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileserver))

	r.HandleFunc("/", Home(app))
	r.HandleFunc("/snippet/view", View(app))
	r.HandleFunc("/snippet/create", Wcreate(app)).Methods(http.MethodPost)
	return r
}
