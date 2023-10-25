package main

// This file stores the things that are performed on the web page

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// this is handler for the home page, this just checks that the page relates to the "/"
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
}

// This is a webpage for vieeing some post results
func (app *application) View(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "We are writing on a view page")
	fmt.Fprintf(w, "We are viewing the snippet with the id: %v", id)

}

func (app *application) WCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Writing something on the WCreate page"))
}
