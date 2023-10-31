package main

import (
	"errors"
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
	validator "github.com/3WDeveloper-GM/webapp-rewrite/internal/validator"
)

type userLoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	validator.Validator
}

func userLogin(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := app.GetTemplateData(r)
		data.Form = userLoginForm{}
		app.Render(w, http.StatusOK, "login.tmpl", data)
	}
}

func userLogoutPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := app.SessionManager.RenewToken(r.Context())
		if err != nil {
			app.ServerError(w, err)
		}

		app.SessionManager.Remove(r.Context(), "authenticatedUserID")

		app.SessionManager.Put(r.Context(), "flash", "You've logged out succesfully")

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func userLoginPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var formdata userLoginForm

		err := app.DecodePostForm(r, &formdata)
		if err != nil {
			app.ClientError(w, http.StatusBadRequest)
		}

		formdata.CheckField(validator.NotBlank(formdata.Email), "email", "This field cannot be blank")
		formdata.CheckField(validator.Matches(formdata.Email, validator.EmailRX), "email", "This must be a valid email address")
		formdata.CheckField(validator.NotBlank(formdata.Password), "password", "This field cannot be blank")

		if !formdata.Valid() {
			data := app.GetTemplateData(r)
			data.Form = formdata
			app.Render(w, http.StatusUnprocessableEntity, "login.tmpl", data)
			return
		}

		id, err := app.Users.Authenticate(formdata.Email, formdata.Password)
		if err != nil {
			if errors.Is(err, models.ErrInvalidCredentials) {
				formdata.AddNonfieldError("Email or password is incorrect")

				data := app.GetTemplateData(r)
				data.Form = formdata
				app.Render(w, http.StatusUnprocessableEntity, "login.tmpl", data)
			} else {
				app.ServerError(w, err)
			}
			return
		}

		err = app.SessionManager.RenewToken(r.Context())
		if err != nil {
			app.ServerError(w, err)
			return
		}

		app.SessionManager.Put(r.Context(), "authenticatedUserID", id)

		http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
	}
}
