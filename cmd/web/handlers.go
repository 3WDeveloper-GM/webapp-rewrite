package main

// This file stores the things that are performed on the web page

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg"
)

// this is handler for the home page, this just checks that the page relates to the "/"
// I use the app *config.Application because it's easier for handling errors between packages
func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			app.NotFound(w)
			return
		}

		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/home.tmpl",
			//"./ui/html/pages/home.bak",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.Errorlog.Println(err.Error())
			//http.Error(w, "Internal Server error", http.StatusInternalServerError)
			app.ServerError(w, err)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.Errorlog.Println(err.Error())
			//http.Error(w, "Internal Server error", http.StatusInternalServerError)
			app.ServerError(w, err)
			return
		}
	}
}

// This is a webpage for vieeing some post results
func View(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "We are writing on a view page")
		fmt.Fprintf(w, "We are viewing the snippet with the id: %v", id)

	}
}

func Wcreate(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {

			w.Header().Set("Allow", http.MethodPost)
			app.ClientError(w, http.StatusMethodNotAllowed)
			//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("Writing something on the WCreate page"))
	}
}
