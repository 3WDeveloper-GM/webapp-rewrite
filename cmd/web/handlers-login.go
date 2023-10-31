package main

import (
	"errors"
	"fmt"
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
	internal "github.com/3WDeveloper-GM/webapp-rewrite/internal/validator"
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

		//Form validation, I start to notice the pattern.
		formdata.CheckField(internal.NotBlank(formdata.Name), "name", "Field cannot be blank")
		formdata.CheckField(internal.NotBlank(formdata.Email), "email", "Field cannot be blank")
		formdata.CheckField(internal.Matches(formdata.Email, internal.EmailRX), "email", "this field must be a valid email address")
		formdata.CheckField(internal.MinChars(formdata.Password, 8), "password", "This field cannot be less than 8 characters")
		formdata.CheckField(internal.NotBlank(formdata.Password), "password", "field cannot be blank")
		formdata.CheckField(internal.MaxChars(formdata.Password, 30), "password", "Field cannot be more than 30 characters.")

		if !formdata.Valid() {
			data := app.GetTemplateData(r)
			data.Form = formdata
			app.Render(w, http.StatusUnprocessableEntity, "signup.tmpl", data)
			return
		}

		err = app.Users.Insert(formdata.Name, formdata.Email, formdata.Password)
		if err != nil {
			if errors.Is(err, models.ErrDuplicateEmail) {
				formdata.AddFieldError("email", "email address already in use")

				data := app.GetTemplateData(r)
				data.Form = formdata
				app.Render(w, http.StatusUnprocessableEntity, "signup.tmpl", data)
				return
			} else {
				app.ServerError(w, err)
			}
			return
		}

		app.SessionManager.Put(r.Context(), "flash", "Your signup was successful, please log in")

		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
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
