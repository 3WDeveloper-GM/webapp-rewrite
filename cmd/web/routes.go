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

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.NotFound(w)
	})

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileserver)).Methods(http.MethodGet)

	dynamicmiddleware := alice.New(app.SessionManager.LoadAndSave, app.CSRFProtectionToken)

	protectedmiddleware := dynamicmiddleware.Append(app.RequireAuthentication)
	//r.HandleFunc("/", Home(app)).Methods(http.MethodGet)
	//r.HandleFunc("/snippet/view/{id}", View(app)).Methods(http.MethodGet)
	//r.HandleFunc("/snippet/create", SnippetCreate(app)).Methods(http.MethodGet)
	//r.HandleFunc("/snippet/create", SnippetPosting(app)).Methods(http.MethodPost)

	//Reworte these using the "gorilla/mux" way. Using the examples of the documentation
	//as a base

	// This are the routes i will use for navigating the webpage
	r.Path("/").Handler(dynamicmiddleware.ThenFunc(Home(app))).Methods(http.MethodGet)
	r.Path("/snippet/view/{id}").Handler(dynamicmiddleware.ThenFunc(View(app))).Methods(http.MethodGet)
	r.Path("/users/signup").Handler(dynamicmiddleware.ThenFunc(userSignup(app))).Methods(http.MethodGet)
	r.Path("/users/signup").Handler(dynamicmiddleware.ThenFunc(userSignupPost(app))).Methods(http.MethodPost)
	r.Path("/users/login").Handler(dynamicmiddleware.ThenFunc(userLogin(app))).Methods(http.MethodGet)
	r.Path("/users/login").Handler(dynamicmiddleware.ThenFunc(userLoginPost(app))).Methods(http.MethodPost)
	// This are the routes I'll use for the authentication part

	r.Path("/users/logout").Handler(protectedmiddleware.ThenFunc(userLogoutPost(app))).Methods(http.MethodPost)
	r.Path("/snippet/create").Handler(protectedmiddleware.ThenFunc(SnippetCreate(app))).Methods(http.MethodGet)
	r.Path("/snippet/create").Handler(protectedmiddleware.ThenFunc(SnippetPosting(app))).Methods(http.MethodPost)

	//Chaining some middleware
	standard := alice.New(app.RecoverPanic, app.LogRequest, config.SecureHeaders)

	return standard.Then(r)
}
