package main

// This file stores the routes that are made when moving on the website.

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileserver))

	app := GetApp()
	r.HandleFunc("/", app.Home)
	r.HandleFunc("/View", app.View)
	r.HandleFunc("/WCreate", app.WCreate)
	return r
}
