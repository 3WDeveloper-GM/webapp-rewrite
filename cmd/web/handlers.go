package main

// This file stores the things that are performed on the web page

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg"
	"github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/templating"
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

		snippets, err := app.Snippets.Latest()
		if err != nil {
			app.ServerError(w, err)
			return
		}

		// This method makes it so when i render a web page I just have to maintain the Render method.
		app.Render(w, http.StatusOK, "home.tmpl", &templating.TemplateData{
			Snippets: snippets,
		})
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

		app.Render(w, http.StatusOK, "view.tmpl", &templating.TemplateData{
			Snippet: snippet,
		})

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
