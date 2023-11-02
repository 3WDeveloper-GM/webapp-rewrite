package main

import (
	"errors"
	"net/http"

	config "github.com/3WDeveloper-GM/webapp-rewrite/cmd/pkg/configuration"
	"github.com/3WDeveloper-GM/webapp-rewrite/internal/models"
	internal "github.com/3WDeveloper-GM/webapp-rewrite/internal/validator"
)

type userpasswordForm struct {
	CurrentPassword       string `form:"currentPassword"`
	NewPassWord           string `form:"newPassword"`
	NewPassWordValidation string `form:"newPasswordValidator"`
	internal.Validator    `form:"-"`
}

func UpdatePasswordPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var formdata userpasswordForm

		err := app.DecodePostForm(r, &formdata)
		if err != nil {
			app.ServerError(w, err)
		}

		formdata.CheckField(internal.NotBlank(formdata.CurrentPassword), "currentPassword", "This field cannot be blank")
		formdata.CheckField(internal.NotBlank(formdata.NewPassWord), "newPassword", "This field cannot be blank")
		formdata.CheckField(internal.NotBlank(formdata.NewPassWordValidation), "newPasswordValidator", "This field cannot be blank")
		formdata.CheckField(internal.MaxChars(formdata.NewPassWord, 30), "newPassword", "Password must be 30 characters or less")
		formdata.CheckField(internal.MinChars(formdata.NewPassWord, 8), "newPassword", "Password must be longer than 8 characters")
		formdata.CheckField(formdata.NewPassWord == formdata.NewPassWordValidation, "newPasswordValidator", "Both passwords must match")

		if !formdata.Valid() {
			data := app.GetTemplateData(r)
			data.Form = formdata
			app.Render(w, http.StatusUnprocessableEntity, "update.tmpl", data)
		}

		userID := app.SessionManager.GetInt(r.Context(), "authenticatedUserID")

		err = app.Users.PasswordUpdate(userID, formdata.CurrentPassword, formdata.NewPassWord)
		if err != nil {
			if errors.Is(err, models.ErrInvalidCredentials) {
				formdata.AddFieldError("currentPassword", "Current password is incorrect")

				data := app.GetTemplateData(r)
				data.Form = formdata
				app.Render(w, http.StatusUnprocessableEntity, "update.tmpl", data)
			} else if err != nil {
				app.ServerError(w, err)
			}
		}

		app.SessionManager.Put(r.Context(), "flash", "You've succesfully changed your password")

		http.Redirect(w, r, "/account/view", http.StatusSeeOther)

	}
}

func UpdatePasswordGet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := app.GetTemplateData(r)
		data.Form = &userpasswordForm{}
		app.Render(w, http.StatusOK, "update.tmpl", data)
	}
}
