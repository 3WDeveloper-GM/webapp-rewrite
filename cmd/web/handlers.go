package main

// This file stores the things that are performed on the web page

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
	"github.com/gorilla/mux"
)

type SnippetPostingValidation struct {
	Title              string `form:"title"`
	Content            string `form:"content"`
	Expires            int    `form:"expires"`
	internal.Validator `form:"-"`
}

// this is handler for the home page, this just checks that the page relates to the "/"
// I use the app *config.Application because it's easier for handling errors between packages
func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.URL.Path != "/" {
		// 	app.NotFound(w)
		// 	return
		// }

		snippets, err := app.Snippets.Latest()
		if err != nil {
			app.ServerError(w, err)
			return
		}

		data := app.CurrentYearTemplateData(r) //Template data that gets the current year
		data.Snippets = snippets               //Getting the snippet data

		// This method makes it so when i render a web page I just have to maintain the Render method.
		app.Render(w, http.StatusOK, "home.tmpl", data)
	}
}

// This is a webpage for vieeing some post results
func View(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
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

		data := app.CurrentYearTemplateData(r)
		data.Snippet = snippet

		app.Render(w, http.StatusOK, "view.tmpl", data)

		// 2023/10/26 yay, i made a functional product at last.
		// Maybe this thing is not that hard after all

	}
}

func SnippetCreate(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := app.CurrentYearTemplateData(r)

		data.Form = &SnippetPostingValidation{
			Expires: 365,
		}

		app.Render(w, http.StatusOK, "create.tmpl", data)
	}
}

func SnippetPosting(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var formdata SnippetPostingValidation

		err := app.DecodePostForm(r, &formdata)
		if err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		err = app.FormDecoder.Decode(&formdata, r.PostForm)
		if err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		formdata.CheckField(internal.NotBlank(formdata.Title), "title", "This field cannot be blank.")
		formdata.CheckField(internal.MaxChars(formdata.Title, 100), "title", "This field cannot contain more than 100 characters.")
		formdata.CheckField(internal.NotBlank(formdata.Content), "content", "This field cannot be blank")
		formdata.CheckField(internal.PermittedInt(formdata.Expires, 1, 7, 365), "expires", "This field can only take values of 1, 7 or 365")

		if !formdata.Valid() {
			data := app.CurrentYearTemplateData(r)
			data.Form = formdata
			app.Render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
			return
		}

		// The SnippetPosting handler doesnt need to check whether the request is
		// a POST or a GET this is done automatically by the gorilla/mux router

		id, err := app.Snippets.Insert(formdata.Title, formdata.Content, formdata.Expires)
		if err != nil {
			fmt.Println("error here")
			app.ServerError(w, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%v", id), http.StatusSeeOther)
	}
}
