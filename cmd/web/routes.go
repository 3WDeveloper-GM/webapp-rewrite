package main

// This file stores the routes that are made when moving on the website.

import (
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileserver))

	app := config.GetApp()
	r.HandleFunc("/", Home(app))
	r.HandleFunc("/View", View(app))
	r.HandleFunc("/WCreate", Wcreate(app))
	return r
}
