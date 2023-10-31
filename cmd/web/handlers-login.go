package main

import (
	"fmt"
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal"
)

type userSignupForm struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	internal.Validator
}

func userSignup(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := app.GetTemplateData(r)
		data.Form = userSignupForm{}
		app.Render(w, http.StatusOK, "signup.tmpl", data)
	}
}

func userSignupPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var formdata userSignupForm

		err := app.DecodePostForm(r, &formdata)
		if err != nil {
			app.ClientError(w, http.StatusBadRequest)
			return
		}

		formdata.CheckField(internal.NotBlank(formdata.Name), "name", "Field cannot be blank")
		formdata.CheckField(internal.NotBlank(formdata.Email), "email", "Field cannot be blank")
		formdata.CheckField(internal.Matches(formdata.Email, internal.EmailRX), "email", "this field must be a valid email address")
		formdata.CheckField(internal.MinChars(formdata.Password, 8), "password", "This field cannot be less than 8 characters")
		formdata.CheckField(internal.NotBlank(formdata.Password), "password", "field cannot be blank")

		if !formdata.Valid() {
			data := app.GetTemplateData(r)
			data.Form = formdata
			app.Render(w, http.StatusOK, "signup.tmpl", data)
		}
	}
}

func userLogin(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Some web form for login")
	}
}

func userLoginPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Some method for login, limited to posting")
	}
}

func userLogoutPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "some method for logout, limited to post")
	}
}
