package main

// This file stores the things that are performed on the web page

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
)

// this is handler for the home page, this just checks that the page relates to the "/"
// I use the app *config.Application because it's easier for handling errors between packages
func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			app.NotFound(w)
			return
		}

		// snippets, err := app.Snippets.Latest()
		// if err != nil {
		// 	app.ServerError(w, err)
		// 	return
		// }

		// for _, snippet := range snippets {
		// 	fmt.Fprintf(w, "%+v\n", snippet)
		// }

		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/home.tmpl",
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
			app.NotFound(w)
			return
		}

		//fmt.Print(id)
		snippet, err := app.Snippets.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.NotFound(w)
			} else {
				app.ServerError(w, err)
			}
			return
		}

		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/view.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ServerError(w, err)
			return
		}

		data := &TemplateData{
			Snippet: snippet,
		}

		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ServerError(w, err)
		}

		// 2023/10/26 yay, i made a functional product at last.
		// Maybe this thing is not that hard after all

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

		title := "O snail"
		content :=
			`
		O snail
		Climb Mount Fuji
		Slowly, but SLowly!

		- Kobayashi Issa
		`
		expires := 7

		id, err := app.Snippets.Insert(title, content, expires)
		if err != nil {
			fmt.Println("error here")
			app.ServerError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
	}
}
